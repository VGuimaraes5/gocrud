package db

import (
	"github.com/VGuimaraes5/gocrud/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	connection *gorm.DB
)

func Init() error {
	dsn := "host=postgres-db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("Error connecting to database: %v", err)
		return err
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		println("Automigration error: %v", err)
		return err
	}

	connection = db
	return nil
}

func GetConnection() *gorm.DB {
	return connection
}
