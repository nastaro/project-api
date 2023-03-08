package database

import (
	"github.com/joho/godotenv"
	"github.com/nastaro/project-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DBK8s")), &gorm.Config{})
	if err != nil {
		db, err = gorm.Open(postgres.Open(os.Getenv("DBLOCAL")), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	}
	db.AutoMigrate(&models.Project{})
	DB = db
}
