package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/app"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/helpers"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/models"
)

type UserController struct {
    App *app.AppContext
}




func (c *UserController) Register (ctx *gin.Context) {
	var userInput models.User
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

		if err := helpers.ValidateUser(userInput); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
	existingUser := models.User{}
	result := c.App.DB.Where("email = ?", userInput.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email is already registered"})
		return
	}

	hashedPassword, err := helpers.HashPassword(userInput.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	
	// Create a new user
	newUser := models.User{
		Username:  userInput.Username,
		Email:     userInput.Email,
		Password:  string(hashedPassword),
	}

	result = c.App.DB.Create(&newUser)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

}

func (c *UserController) Login(ctx *gin.Context) {
	var loginInput app.LoginInput
	if err := ctx.ShouldBindJSON(&loginInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

		if err := helpers.ValidateLoginInput(loginInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := c.App.DB.Where("email = ?", loginInput.Email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if match := helpers.CheckPasswordHash(loginInput.Password, user.Password); match != true {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	secretKey := []byte("your_secret_key") 
	token, err := helpers.GenerateJWT(user.ID, secretKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}