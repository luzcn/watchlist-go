package db

import (
	"time"
)

// Product a testing db nodel
type Product struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Price     string    `gorm:"not null" json:"price"`
	Notes     []Note    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}

// Note the db table of case watch list
type Note struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	ProductId uint      `json:"product_id"`
	Note      string    `gorm:"not null" json:"note"`
	Context   string    `gorm:"not null" json:"context"`
	CreatedAt time.Time `json:"created_at"`
}
