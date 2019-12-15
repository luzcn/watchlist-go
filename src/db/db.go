package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type Env struct {
	DB *gorm.DB
}

func (env *Env) AddNote(note *Notes) {
	env.DB.Create(note)
}

// get the first db record
func (env *Env) GetNote(note *Notes) {
	if err := env.DB.Take(note).Error; err != nil {
		log.Println(err)
	}
}

func (env *Env) ListNotes(notes *[]Notes) {

	if err := env.DB.Find(&notes).Error; err != nil {
		log.Println(err)
	}

}
