package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// Setup user routes
	setupUserRoutes(router)

	// Setup photo routes

	// Add more route setups as needed
}

func setupUserRoutes(router *gin.Engine){
	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.CreateUser)
	}
}