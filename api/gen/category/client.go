// Code generated by goa v3.1.3, DO NOT EDIT.
//
// category client
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package category

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "category" service client.
type Client struct {
	ListEndpoint goa.Endpoint
}

// NewClient initializes a "category" service client given the endpoints.
func NewClient(list goa.Endpoint) *Client {
	return &Client{
		ListEndpoint: list,
	}
}

// List calls the "list" endpoint of the "category" service.
func (c *Client) List(ctx context.Context) (res []*Category, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*Category), nil
}
