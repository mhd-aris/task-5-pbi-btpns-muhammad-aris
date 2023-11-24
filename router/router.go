package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/app"
)

func SetupRoutes(myApp *app.App) *gin.Engine {
	r := gin.Default()

	userRoutes := NewUserRoutes(myApp)
	userRoutes.SetupRoutes(r.Group("/users"))


	return r
}
