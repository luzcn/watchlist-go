package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Product a testing db nodel
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// Notes the db table of case watch list
type Notes struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Note      string     `gorm:"not null" json:"note"`
	Context   string     `gorm:"not null" json:"context"`
}
