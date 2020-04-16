package hub

import (
	"context"
	"log"

	resource "github.com/tektoncd/hub/api/gen/resource"
)

// resource service example implementation.
// The example methods log the requests and return zero values.
type resourcesrvc struct {
	logger *log.Logger
}

// NewResource returns the resource service implementation.
func NewResource(logger *log.Logger) resource.Service {
	return &resourcesrvc{logger}
}

// Get all tasks and pipelines.
func (s *resourcesrvc) All(ctx context.Context, p *resource.AllPayload) (res resource.ResourceCollection, err error) {
	s.logger.Print("resource.All")
	return
}

// Get all versions of a resource
func (s *resourcesrvc) AllVersions(ctx context.Context, p *resource.AllVersionsPayload) (res []*resource.Resource, err error) {
	s.logger.Print("resource.AllVersions")
	return
}

// Get User's rating of a resource
func (s *resourcesrvc) GetRating(ctx context.Context, p *resource.GetRatingPayload) (res *resource.ResourceRating, err error) {
	res = &resource.ResourceRating{}
	s.logger.Print("resource.GetRating")
	return
}

// Update User's rating of a resource
func (s *resourcesrvc) UpdateRating(ctx context.Context, p *resource.UpdateRatingPayload) (err error) {
	s.logger.Print("resource.UpdateRating")
	return
}
