package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Env struct {
	DB *gorm.DB
}

type DataAccess interface {
	CreateProduct(*Product)
	GetProduct() Product
}

// We can use NewRecord or Create to insert new row in DB
// NewRecord return true/false
// Create can return error message
func (env *Env) CreateProduct(p *Product) {
	if err := env.DB.Create(p).Error; err != nil {
		panic(err)
	}
}

// get the first db record
func (env *Env) GetProduct() Product {
	p := Product{}
	env.DB.Take(&p)

	env.DB.Model(p).Related(&p.Notes)

	return p
}
