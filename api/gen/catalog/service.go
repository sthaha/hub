// Code generated by goa v3.2.2, DO NOT EDIT.
//
// catalog service
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package catalog

import (
	"context"

	catalogviews "github.com/tektoncd/hub/api/gen/catalog/views"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// The Catalog Service exposes endpoints to interact with catalogs
type Service interface {
	// Refreshes Tekton Catalog
	Refresh(context.Context, *RefreshPayload) (res *Job, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "catalog"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"Refresh"}

// RefreshPayload is the payload type of the catalog service Refresh method.
type RefreshPayload struct {
	// JWT
	Token string
}

// Job is the result type of the catalog service Refresh method.
type Job struct {
	// id of the job
	ID uint
	// status of the job
	Status string
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal-error",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not-found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewJob initializes result type Job from viewed result type Job.
func NewJob(vres *catalogviews.Job) *Job {
	return newJob(vres.Projected)
}

// NewViewedJob initializes viewed result type Job from result type Job using
// the given view.
func NewViewedJob(res *Job, view string) *catalogviews.Job {
	p := newJobView(res)
	return &catalogviews.Job{Projected: p, View: "default"}
}

// newJob converts projected type Job to service type Job.
func newJob(vres *catalogviews.JobView) *Job {
	res := &Job{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	return res
}

// newJobView projects result type Job to projected type JobView using the
// "default" view.
func newJobView(res *Job) *catalogviews.JobView {
	vres := &catalogviews.JobView{
		ID:     &res.ID,
		Status: &res.Status,
	}
	return vres
}
