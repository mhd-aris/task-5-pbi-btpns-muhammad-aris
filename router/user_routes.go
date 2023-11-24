package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/app"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/controllers"
)


type UserRoutes struct{
	App *app.App
}

func NewUserRoutes(myApp *app.App) *UserRoutes{
	return &UserRoutes{
		App: myApp,
	}
}

func (us *UserRoutes) SetupRoutes(group *gin.RouterGroup){
	userController := controllers.NewUserController(us.App)
	group.POST("/register", userController.CreateUser)
	group.GET("/", userController.GetUser)
}