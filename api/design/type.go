package design

import (
	. "goa.design/goa/v3/dsl"
)

var Resource = ResultType("application/vnd.hub.resource", func() {
	Description("Describes a resource retrieved from catalog")
	TypeName("Resource")

	Attribute("id", UInt, "ID is the unique id of the resource", func() {
		Example("id", 1)
	})
	Attribute("name", String, "Name of the resource", func() {
		Example("name", "buildah")
	})
	Attribute("type", String, "Type of the resource", func() {
		Example("type", "Task")
	})
	Attribute("catalog", Catalog, "Catalog to which resource belongs to")
	Attribute("description", String, "Description of the resource", func() {
		Example("description", "Buildah is a build tool to create images from source code.")
	})
	Attribute("versions", ArrayOf(ResourceVersion), "Different Versions of the resource")
	Attribute("tags", ArrayOf(ResourceTag), "Tags associated to the resource")
	Attribute("rating", Float64, "Rating of resource", func() {
		Example("rating", 2.3)
		Minimum(0)
		Maximum(5)
	})
	Attribute("last_updated_at", String, "TimeStamp the resource last updated at", func() {
		Example("last_updated_at", "1232342423")
	})
	Required("id", "name", "catalog", "description", "type", "versions", "tags", "rating", "last_updated_at")
})

var Catalog = Type("Catalog", func() {
	Attribute("id", UInt, "ID is the unique id of the catalog", func() {
		Example("id", 1)
	})
	Attribute("type", String, "Type of catalog", func() {
		Example("type", "Official")
	})
	Required("id", "type")
})

var ResourceVersion = Type("ResourceVersion", func() {
	Attribute("id", UInt, "ID is the unique id of the version", func() {
		Example("id", 1)
	})
	Attribute("version", String, "Version of resource", func() {
		Example("version", "2.3")
	})
	Required("id", "version")
})

var ResourceTag = Type("ResourceTag", func() {
	TypeName("Tag")
	Attribute("id", UInt, "ID is the unique id of the tag", func() {
		Example("id", 1)
	})
	Attribute("name", String, "Name of the tag", func() {
		Example("name", "notification")
	})
	Required("id", "name")
})

var ResourceRating = Type("ResourceRating", func() {
	Attribute("rating", UInt, "Rating of the resource", func() {
		Example("rating", 3)
		Minimum(0)
		Maximum(5)
	})
	Required("rating")
})

var Category = Type("Category", func() {
	Attribute("id", UInt, "ID is the unique id of the category", func() {
		Example("id", 1)
	})
	Attribute("name", String, "Name of the Category", func() {
		Example("name", "Notification")
	})
	Attribute("tags", ArrayOf(ResourceTag), "Tags associated with the category")
	Required("id", "name", "tags")
})

var ResourceVersionDetail = Type("ResourceVersionDetails", func() {

	Attribute("version", String, "Version of the resource", func() {
		Example("version", "1.0")
	})
	Attribute("description", String, "Description of the resource's version", func() {
		Example("description", "Buildah is a build tool to create images from source code.")
	})
	Attribute("url", String, "Url of the resource's version", func() {
		Example("url", "https://github.com/tektoncd/catalog/tasks/buildah/1.0")
	})
	Required("version", "description", "url")
})
