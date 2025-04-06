package routers

import (
	"github.com/gin-gonic/gin"
)

func setUpRouterPing(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "bidang",
		})
	})
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	setUpRouterPing(api)
	
}
