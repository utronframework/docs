package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Page struct {
	gorm.Model
	Title    string
	Body     string
	Sections []Section `gorm:"sections"`
}

type Section struct {
	gorm.Model
	ParentID uint
	Name     string
	PageID   uint
}
