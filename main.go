package main

import (
	"myapi/internal/bootstrap"
)

// @title MyAPI - API Docs
// @version 2.0
// @description An API Documentation for MyAPI
// @termsOfService https://myapi.com/term-of-service

// @contact.name API Support
// @contact.url https://myapi.com/en/contactus
// @contact.email info@myapi.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey 	ApiKeyAuth
// @in                          header
// @name                        X-API-KEY
// @host myapi
// @BasePath /
// @schemes http
func main() {
  app := bootstrap.NewApp()
  app.Run()
}