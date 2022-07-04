package main

import (
	"github.com/david-luk4s/api-go-gin/database"
	"github.com/david-luk4s/api-go-gin/routers"
)

// gin-swagger middleware
// swagger embed files
// gin-swagger middleware
// swagger embed files
// @title           API for Students
// @version         1.0
// @description     This is a sample api for Students
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.github    https://github.com/david-luk4s/api-go-gin
// @contact.username  david-luk4s

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	database.ConnectionDB()
	routers.HandleRequests()
}
