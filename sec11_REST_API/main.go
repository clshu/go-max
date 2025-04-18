package main

import (
	"example.com/restapi/db"
	"example.com/restapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")

}
