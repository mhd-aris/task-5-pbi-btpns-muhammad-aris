package database

import (
	"fmt"

	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/config"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func NewDB(dbConfig *config.DatabaseConfig) *gorm.DB{

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require TimeZone=Asia/Jakarta",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
        panic("Failed to connect to database")
    }

	err = db.Debug().AutoMigrate(&models.User{}, &models.Photo{})
	if err != nil {
        panic("Failed to migrate database")
    }

	return db
}