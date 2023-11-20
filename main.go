package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/app"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/config"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/models"
)

func init(){
	config.Load()
}

func main(){
	r := gin.Default()

	db := app.InitDB()
	r.GET("/ping", func(c * gin.Context){
			
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})
	r.POST("/pong", func(c * gin.Context){
		var user models.User

		user.Username= "Test"
		user.Email= "test@mail"
		user.Password= "123"
		
		result := db.Create(&user)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   result.Error.Error(),
				"message": "Failed to create user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
			"data": user,
		})
	})

	r.Run()
}