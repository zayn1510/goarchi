package main

import (
	"github.com/zayn1510/goarchi/config"
	"log"
)

func main() {
	config.ConnectDB()
	db := config.GetDB()
	if db == nil {
		log.Fatal("Database tidak terkoneksi!")
	}
}
