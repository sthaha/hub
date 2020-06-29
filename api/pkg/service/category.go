package hub

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"go.uber.org/zap"

	category "github.com/tektoncd/hub/api/gen/category"
	app "github.com/tektoncd/hub/api/pkg/app"
)

// category service example implementation.
// The example methods log the requests and return zero values.
type categorysrvc struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

// NewCategory returns the category service implementation.
func NewCategory(api *app.ApiConfig) category.Service {
	return &categorysrvc{api.Logger(), api.DB()}
}

// Get all Categories with their tags sorted by name
func (s *categorysrvc) All(ctx context.Context) (res []*category.Category, err error) {
	var all []model.Category
	if err := s.db.Order("name").Preload("Tags").Find(&all).Error; err != nil {
		s.logger.Error(err)
		return []*category.Category{}, category.MakeInternalError(fmt.Errorf("Failed to fetch categories"))
	}

	for _, c := range all {
		tags := []*category.Tag{}
		for _, t := range c.Tags {
			tags = append(tags, &category.Tag{
				ID:   t.ID,
				Name: t.Name,
			})
		}
		res = append(res, &category.Category{
			ID:   c.ID,
			Name: c.Name,
			Tags: tags,
		})
	}

	return res, nil
}
