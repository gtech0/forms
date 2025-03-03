package main

import (
	_ "hedgehog-forms/docs"
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

//	@title		Forms API
// @version         0.01
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8001
// @BasePath  /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	router := web.NewRouter()
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
