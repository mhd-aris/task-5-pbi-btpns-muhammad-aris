package routers

import (
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/app"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/controllers"
)


type UserRoutes struct{
	App *app.AppContext
}

func InitUserRoutes(appCtx *app.AppContext) {
	userController := controllers.UserController{App: appCtx}
	
	userRoutes := appCtx.Router.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
	}
}