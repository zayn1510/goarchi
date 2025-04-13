package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zayn1510/goarchi/routers"
)

func main() {
	router := gin.Default()
	routers.RegisterRoutes(router)
	router.Run(":8082")
}
