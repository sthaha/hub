package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("category", func() {
	Description("The category service gives categories details")

	//Method to get all categories with their tags
	Method("Categories", func() {
		Description("Get all Categories with their associated tags.")
		Result(func() {
			Attribute("data", ArrayOf(Category))
			Attribute("errors", ArrayOf(String)) // To be updated
		})
		HTTP(func() {
			GET("/categories")
			Response(StatusOK)
		})
	})
})