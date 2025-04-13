package controllers

import "github.com/gin-gonic/gin"

type UsersController struct{}

func (ctrl *UsersController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from UsersController",
	})
}
