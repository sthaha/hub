package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("resource", func() {
	Description("The resource service gives resource details")

	//Method to get all resources
	Method("All", func() {
		Description("Get all tasks and pipelines.")
		Payload(func() {
			Attribute("limit", UInt, "Number of Resource to fetch", func() {
				Default(100)
			})
		})
		Result(CollectionOf(Resource))
		HTTP(func() {
			GET("/resources")
			Param("limit")
			Response(StatusOK)
		})
	})

	//Method to get all versions of a resource
	Method("AllVersions", func() {
		Description("Get all versions of a resource")
		Payload(func() {
			Attribute("resourceID", UInt, "Id of the Resource")
		})
		Result(ArrayOf(Resource))
		HTTP(func() {
			GET("/resource/{resourceID}/versions")
			Response(StatusOK)
		})
	})

	//Method to get user's rating of a resource
	Method("GetRating", func() {
		Description("Get User's rating of a resource")
		Payload(func() {
			Attribute("resourceID", UInt, "Id of the Resource")
			Required("resourceID")
		})
		Result(ResourceRating)
		HTTP(func() {
			GET("/resource/{resourceID}/rating")
			Response(StatusOK)
		})
	})

	//Method to update user's rating of a resource
	Method("UpdateRating", func() {
		Description("Update User's rating of a resource")
		Payload(func() {
			Attribute("rating", UInt, "Rating of resource to be updated")
			Attribute("resourceID", UInt, "Id of the Resource")
			Required("rating", "resourceID")
		})
		HTTP(func() {
			PUT("/resource/{resourceID}/rating")
			Response(StatusOK)
		})
	})
})
