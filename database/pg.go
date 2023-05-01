package database

import (
	"log"

	"github.com/Group-8-H8/fp-3/config"
	"github.com/Group-8-H8/fp-3/entity"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open(config.GetDBConfig())
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err = db.AutoMigrate(&entity.User{}, &entity.Category{}, &entity.Task{}); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Connected to DB!")
}

func GetPostgresInstance() *gorm.DB {
	return db
}
