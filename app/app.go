package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppContext struct {
	DB *gorm.DB
	Router *gin.Engine
}

type LoginInput struct{
	Email string
	Password string
}

func New(db *gorm.DB) *AppContext {
	return &AppContext{
		DB:     db,
		Router: gin.New(),
	}
}



