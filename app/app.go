package app

import (
	"fmt"

	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/config"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}



func NewApp() *App{
	db := initDB()
	
	return &App{DB: db}
}


func initDB() *gorm.DB{
	dbConfig := config.GetDatabaseConfig()

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

