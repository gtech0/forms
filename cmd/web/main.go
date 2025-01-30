package main

import (
	"hedgehog-forms/internal/core/enviroment"
	"hedgehog-forms/internal/core/file"
	"hedgehog-forms/internal/ui/web"
	"hedgehog-forms/pkg/database"
	"hedgehog-forms/pkg/migration"
	"log"
)

func init() {
	enviroment.Load()
	database.Connect()
	migration.Sync()
	file.InitializeClient()
}

func main() {
	router := web.NewRouter()
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
