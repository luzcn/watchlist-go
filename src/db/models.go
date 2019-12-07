package db

import "github.com/jinzhu/gorm"

// Product a testing db nodel
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// Notes the db table of case watch list
type Notes struct {
	gorm.Model
	Note    string `gorm:"not null"`
	Context string `gorm:"not null"`
}
