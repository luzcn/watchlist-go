package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Env struct {
	DB *gorm.DB
}

func (env *Env) AddNote(note *Notes) {
	env.DB.Create(note)
}
