package main

import (
	"apidanadesa/app/controllers"
	"apidanadesa/app/middleware"
	"apidanadesa/routers"
	"github.com/gin-gonic/gin"
)

func SetUpRouterBidang(rg *gin.RouterGroup) {
	c := controllers.NewControllerBidang()
	bidang := rg.Group("/bidang")
	bidang.GET("/", c.GetAllBidangs)
}
func main() {
	router := gin.Default()
	middleware.SetCors(router)
	router.Use(middleware.ErrorHandler())
	routers.RegisterRoutes(router)
	router.Run(":8080")
}
