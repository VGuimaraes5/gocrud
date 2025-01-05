package postgres

import (
	"fmt"

	"github.com/VGuimaraes5/gocrud/configs"
	"github.com/VGuimaraes5/gocrud/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		configs.Database.Host,
		configs.Database.User,
		configs.Database.Password,
		configs.Database.DBName,
		configs.Database.Port,
		configs.Database.SSLMode,
		configs.Database.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("Error connecting to database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		println("Automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
