package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/app"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/models"
)

type UserController struct{
	App *app.App
}

func NewUserController(myApp *app.App) *UserController {
	return &UserController{
		App: myApp,
	}
}


func (c *UserController) CreateUser(ctx *gin.Context) {
	var userInput models.User
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	result := c.App.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)

}

func (c *UserController) GetUser(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "Hello",
	})
}