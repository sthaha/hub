package hub

import (
	"context"
	"log"

	category "github.com/tektoncd/hub/api/gen/category"
)

// category service example implementation.
// The example methods log the requests and return zero values.
type categorysrvc struct {
	logger *log.Logger
}

// NewCategory returns the category service implementation.
func NewCategory(logger *log.Logger) category.Service {
	return &categorysrvc{logger}
}

// Get all Categories with their associated tags.
func (s *categorysrvc) Categories(ctx context.Context) (res *category.CategoriesResult, err error) {
	res = &category.CategoriesResult{}
	s.logger.Print("category.Categories")
	return
}
