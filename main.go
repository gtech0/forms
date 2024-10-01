package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hedgehog-forms/database"
	"hedgehog-forms/enviroment"
	"log"
)

func init() {
	enviroment.Load()
	database.Connect()
	database.Sync()
}

func main() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	router.Use(cors.New(corsConfig))

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
