package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
	"todoapp/models"
)

func Connect() *gorm.DB {

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOSTNAME")
	dbType := os.Getenv("DB_TYPE")

	if dbHost == "" || dbName == "" {
		log.Println(".env file is empty. Please add the environment variables in .env file first and run the server again!")
	}

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string

	db, err := gorm.Open(dbType, dbUri)
	if err != nil {
		panic(err)
	}
	db.Debug().AutoMigrate(models.Todo{})
	return db

}
