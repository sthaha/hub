package model

import (
	"github.com/jinzhu/gorm"
)

type (

	//Category model represents categories which associated with group of tags
	Category struct {
		gorm.Model
		Name string `gorm:"size:100;not null;unique"`
		Tags []Tag
	}

	// Tag model represents tags associated with a resource
	Tag struct {
		gorm.Model
		Name       string `gorm:"size:100;not null;unique"`
		Category   Category
		CategoryID int
	}
)
