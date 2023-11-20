package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhd-aris/task-5-pbi-btpns-muhammad-aris/models"
)

func CreateUser(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success" : false,
			"message" : "Request body not valid!",
			"data" : nil,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success" : true,
		"message" : "User created!",
		"data" : user,
	})
}