package main

import (
	"github.com/david-luk4s/api-go-gin/database"
	"github.com/david-luk4s/api-go-gin/routers"
)

func main() {
	database.ConnectionDB()
	routers.HandleRequests()
}
