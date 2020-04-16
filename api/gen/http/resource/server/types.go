// Code generated by goa v3.1.1, DO NOT EDIT.
//
// resource HTTP server types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	resource "github.com/tektoncd/hub/api/gen/resource"
	resourceviews "github.com/tektoncd/hub/api/gen/resource/views"
	goa "goa.design/goa/v3/pkg"
)

// UpdateRatingRequestBody is the type of the "resource" service "UpdateRating"
// endpoint HTTP request body.
type UpdateRatingRequestBody struct {
	// Rating of resource to be updated
	Rating *uint `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// ResourceResponseCollection is the type of the "resource" service "All"
// endpoint HTTP response body.
type ResourceResponseCollection []*ResourceResponse

// AllVersionsResponseBody is the type of the "resource" service "AllVersions"
// endpoint HTTP response body.
type AllVersionsResponseBody []*ResourceResponse

// GetRatingResponseBody is the type of the "resource" service "GetRating"
// endpoint HTTP response body.
type GetRatingResponseBody struct {
	// Rating of the resource
	Rating uint `form:"rating" json:"rating" xml:"rating"`
}

// ResourceResponse is used to define fields on response body types.
type ResourceResponse struct {
	// ID is the unique id of the resource
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of the resource
	Name string `form:"name" json:"name" xml:"name"`
	// Type of the resource
	Type string `form:"type" json:"type" xml:"type"`
	// Catalog to which resource belongs to
	Catalog *CatalogResponse `form:"catalog" json:"catalog" xml:"catalog"`
	// Description of the resource
	Description string `form:"description" json:"description" xml:"description"`
	// Different Versions of the resource
	Versions []*ResourceVersionResponse `form:"versions" json:"versions" xml:"versions"`
	// Tags associated to the resource
	Tags []*Tag `form:"tags" json:"tags" xml:"tags"`
	// Rating of resource
	Rating float64 `form:"rating" json:"rating" xml:"rating"`
	// TimeStamp the resource last updated at
	LastUpdatedAt string `form:"last_updated_at" json:"last_updated_at" xml:"last_updated_at"`
}

// CatalogResponse is used to define fields on response body types.
type CatalogResponse struct {
	// ID is the unique id of the catalog
	ID uint `form:"id" json:"id" xml:"id"`
	// Type of catalog
	Type string `form:"type" json:"type" xml:"type"`
}

// ResourceVersionResponse is used to define fields on response body types.
type ResourceVersionResponse struct {
	// ID is the unique id of the version
	ID uint `form:"id" json:"id" xml:"id"`
	// Version of resource
	Version string `form:"version" json:"version" xml:"version"`
}

// Tag is used to define fields on response body types.
type Tag struct {
	// ID is the unique id of the tag
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of the tag
	Name string `form:"name" json:"name" xml:"name"`
}

// NewResourceResponseCollection builds the HTTP response body from the result
// of the "All" endpoint of the "resource" service.
func NewResourceResponseCollection(res resourceviews.ResourceCollectionView) ResourceResponseCollection {
	body := make([]*ResourceResponse, len(res))
	for i, val := range res {
		body[i] = marshalResourceviewsResourceViewToResourceResponse(val)
	}
	return body
}

// NewAllVersionsResponseBody builds the HTTP response body from the result of
// the "AllVersions" endpoint of the "resource" service.
func NewAllVersionsResponseBody(res []*resource.Resource) AllVersionsResponseBody {
	body := make([]*ResourceResponse, len(res))
	for i, val := range res {
		body[i] = marshalResourceResourceToResourceResponse(val)
	}
	return body
}

// NewGetRatingResponseBody builds the HTTP response body from the result of
// the "GetRating" endpoint of the "resource" service.
func NewGetRatingResponseBody(res *resource.ResourceRating) *GetRatingResponseBody {
	body := &GetRatingResponseBody{
		Rating: res.Rating,
	}
	return body
}

// NewAllPayload builds a resource service All endpoint payload.
func NewAllPayload(limit uint) *resource.AllPayload {
	v := &resource.AllPayload{}
	v.Limit = limit

	return v
}

// NewAllVersionsPayload builds a resource service AllVersions endpoint payload.
func NewAllVersionsPayload(resourceID uint) *resource.AllVersionsPayload {
	v := &resource.AllVersionsPayload{}
	v.ResourceID = &resourceID

	return v
}

// NewGetRatingPayload builds a resource service GetRating endpoint payload.
func NewGetRatingPayload(resourceID uint) *resource.GetRatingPayload {
	v := &resource.GetRatingPayload{}
	v.ResourceID = resourceID

	return v
}

// NewUpdateRatingPayload builds a resource service UpdateRating endpoint
// payload.
func NewUpdateRatingPayload(body *UpdateRatingRequestBody, resourceID uint) *resource.UpdateRatingPayload {
	v := &resource.UpdateRatingPayload{
		Rating: *body.Rating,
	}
	v.ResourceID = resourceID

	return v
}

// ValidateUpdateRatingRequestBody runs the validations defined on
// UpdateRatingRequestBody
func ValidateUpdateRatingRequestBody(body *UpdateRatingRequestBody) (err error) {
	if body.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rating", "body"))
	}
	return
}
