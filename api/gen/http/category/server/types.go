// Code generated by goa v3.1.1, DO NOT EDIT.
//
// category HTTP server types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	category "github.com/tektoncd/hub/api/gen/category"
)

// CategoriesResponseBody is the type of the "category" service "Categories"
// endpoint HTTP response body.
type CategoriesResponseBody struct {
	Data   []*CategoryResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Errors []string                `form:"errors,omitempty" json:"errors,omitempty" xml:"errors,omitempty"`
}

// CategoryResponseBody is used to define fields on response body types.
type CategoryResponseBody struct {
	// ID is the unique id of the category
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of the Category
	Name string `form:"name" json:"name" xml:"name"`
	// Tags associated with the category
	Tags []*Tag `form:"tags" json:"tags" xml:"tags"`
}

// Tag is used to define fields on response body types.
type Tag struct {
	// ID is the unique id of the tag
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of the tag
	Name string `form:"name" json:"name" xml:"name"`
}

// NewCategoriesResponseBody builds the HTTP response body from the result of
// the "Categories" endpoint of the "category" service.
func NewCategoriesResponseBody(res *category.CategoriesResult) *CategoriesResponseBody {
	body := &CategoriesResponseBody{}
	if res.Data != nil {
		body.Data = make([]*CategoryResponseBody, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalCategoryCategoryToCategoryResponseBody(val)
		}
	}
	if res.Errors != nil {
		body.Errors = make([]string, len(res.Errors))
		for i, val := range res.Errors {
			body.Errors[i] = val
		}
	}
	return body
}
