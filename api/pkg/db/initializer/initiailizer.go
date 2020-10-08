package initializer

import (
	"context"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tektoncd/hub/api/gen/log"
	"github.com/tektoncd/hub/api/pkg/app"
	"github.com/tektoncd/hub/api/pkg/db/model"
)

// Initializer defines the configuration required for initailizer
// to populate the tables
type Initializer struct {
	app.Service
	db   *gorm.DB
	log  *log.Logger
	data *app.Data
}

// New returns the Initializer implementation.
func New(ctx context.Context, api app.BaseConfig) *Initializer {
	service := api.Service("initiailizer")
	return &Initializer{
		Service: service,
		db:      service.DB(ctx),
		log:     service.Logger(ctx),
		data:    api.Data(),
	}
}

// Run executes the func which populate the tables
func (i *Initializer) Run() error {

	config := model.Config{}
	if err := i.db.Model(&config).FirstOrInit(&config).Error; err != nil {
		i.log.Error(err)
		return err
	}

	i.log.Infof("%v", config)

	if config.Checksum == i.data.Checksum {
		i.log.Info("SKIP: Config refresh as config file has not changed")
		return nil
	}

	updateConfig := func(db *gorm.DB) error {
		// Updates the config checksum
		config.Checksum = i.data.Checksum
		if err := db.Save(&config).Error; err != nil {
			i.log.Error(err)
			return err
		}
		return nil
	}

	return withTransaction(i.db,
		i.addCategories,
		i.addCatalogs,
		i.addUsers,
		updateConfig,
	)
}

func (i *Initializer) addCategories(db *gorm.DB) error {

	for _, c := range i.data.Categories {
		cat := &model.Category{Name: c.Name}
		if err := db.Where(cat).FirstOrCreate(cat).Error; err != nil {
			i.log.Error(err)
			return err
		}
		for _, t := range c.Tags {
			tag := &model.Tag{Name: t, CategoryID: cat.ID}
			if err := db.Where(tag).FirstOrCreate(tag).Error; err != nil {
				i.log.Error(err)
				return err
			}
		}
	}
	return nil
}

func (i *Initializer) addCatalogs(db *gorm.DB) error {

	for _, c := range i.data.Catalogs {
		cat := &model.Catalog{
			Name:       c.Name,
			Org:        c.Org,
			Type:       c.Type,
			URL:        c.URL,
			Revision:   c.Revision,
			ContextDir: c.ContextDir,
		}
		if err := db.Where(&model.Catalog{Name: c.Name, Org: c.Org}).FirstOrCreate(cat).Error; err != nil {
			i.log.Error(err)
			return err
		}
	}
	return nil
}

func (i *Initializer) addUsers(db *gorm.DB) error {

	for _, s := range i.data.Scopes {

		// check if scopes exist or create it
		q := db.Where(&model.Scope{Name: s.Name})

		scope := &model.Scope{}
		if err := q.FirstOrCreate(&scope).Error; err != nil {
			i.log.Error(err)
			return err
		}

		for _, userID := range s.Users {

			// Checks if user exists
			q := db.Where("LOWER(github_login) = ?", strings.ToLower(userID))

			user := &model.User{}
			if err := q.First(&user).Error; err != nil {
				// If user not found then log and continue
				if gorm.IsRecordNotFoundError(err) {
					i.log.Warnf("user %s not found: %s", userID, err)
					continue
				}
				i.log.Error(err)
				return err
			}

			// Add scopes for user if not added already
			us := model.UserScope{UserID: user.ID, ScopeID: scope.ID}
			q = db.Model(&model.UserScope{}).Where(&us)

			if err := q.FirstOrCreate(&us).Error; err != nil {
				i.log.Error(err)
				return err
			}
		}

	}
	return nil
}

func withTransaction(db *gorm.DB, fns ...func(*gorm.DB) error) error {
	txn := db.Begin()
	for _, fn := range fns {
		if err := fn(txn); err != nil {
			txn.Rollback()
			return err
		}
	}

	txn.Commit()
	return nil
}
