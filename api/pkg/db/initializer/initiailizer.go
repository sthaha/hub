package initializer

import (
	"context"
	"fmt"
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

	i.db = i.db.Begin()

	// Check if config data is changed
	var config *model.Config
	var err error
	if config, err = i.alreadyApplied(); err != nil {
		return err
	}

	if err := i.addCategories(); err != nil {
		i.db.Rollback()
		return err
	}
	if err := i.addCatalogs(); err != nil {
		i.db.Rollback()
		return err
	}
	if err := i.addUsers(); err != nil {
		i.db.Rollback()
		return err
	}

	// Updates the config checksum
	if config.Checksum != i.data.Checksum {
		config.Checksum = i.data.Checksum
		if err := i.db.Save(config).Error; err != nil {
			i.log.Error(err)
			return err
		}
	}

	i.db.Commit()
	return nil
}

func (i *Initializer) alreadyApplied() (*model.Config, error) {

	db := i.db

	// Checks if table exists
	if !db.HasTable(&model.Config{}) {
		return nil, fmt.Errorf("config table not found")
	}

	// Check if there is an entry in Config table for checksum
	config := &model.Config{}
	if err := db.First(&config).Error; err != nil {

		// If there is no entry then create a new checksum entry
		if gorm.IsRecordNotFoundError(err) {
			newConfig := &model.Config{Checksum: i.data.Checksum}
			if err := db.Create(newConfig).Error; err != nil {
				i.log.Error(err)
				return nil, err
			}
			return newConfig, nil
		}
		i.log.Error(err)
		return nil, err
	}
	if config.Checksum == i.data.Checksum {
		i.log.Info("SKIP: Config refresh as config file has not changed")
		return nil, fmt.Errorf("skip-config-refresh")
	}
	return config, nil
}

func (i *Initializer) addCategories() error {

	db := i.db

	// Checks if tables exists
	if !db.HasTable(&model.Category{}) || !db.HasTable(model.Tag{}) {
		return fmt.Errorf("categories or tags table not found")
	}

	for _, c := range i.data.Categories {
		cat := &model.Category{Name: c.Name}
		if err := db.Where(cat).FirstOrCreate(cat).
			Error; err != nil {
			i.log.Error(err)
			return err
		}
		for _, t := range c.Tags {
			tag := &model.Tag{Name: t, CategoryID: cat.ID}
			if err := db.Where(tag).FirstOrCreate(tag).
				Error; err != nil {
				i.log.Error(err)
				return err
			}
		}
	}
	return nil
}

func (i *Initializer) addCatalogs() error {

	db := i.db

	// Checks if tables exists
	if !db.HasTable(&model.Catalog{}) {
		return fmt.Errorf("catalogs table not found")
	}

	for _, c := range i.data.Catalogs {
		cat := &model.Catalog{
			Name:       c.Name,
			Org:        c.Org,
			Type:       c.Type,
			URL:        c.URL,
			Revision:   c.Revision,
			ContextDir: c.ContextDir,
		}
		if err := db.Where(&model.Catalog{Name: c.Name, Org: c.Org}).
			FirstOrCreate(cat).
			Error; err != nil {
			i.log.Error(err)
			return err
		}
	}
	return nil
}

func (i *Initializer) addUsers() error {

	db := i.db

	// Checks if tables exists
	if !db.HasTable(&model.User{}) || !db.HasTable(model.Scope{}) {
		return fmt.Errorf("user or scope table not found")
	}

	for _, s := range i.data.Scopes {

		// Check if scopes exist
		q := db.Where(&model.Scope{Name: s.Name})

		scope := &model.Scope{}
		if err := q.First(&scope).Error; err != nil {

			// If scope does not exist then return
			if gorm.IsRecordNotFoundError(err) {
				i.log.Errorf("scope (%s) does not exist: %s", s.Name, err)
				return fmt.Errorf("invalid-scope")
			}
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
					i.log.Errorf("user %s not found: %s", userID, err)
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
