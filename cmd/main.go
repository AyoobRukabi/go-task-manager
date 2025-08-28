package main

import (
	"github.com/gin-gonic/gin"
	"github.com/AyoobRukabi/go-task-manager/db"
	"github.com/AyoobRukabi/go-task-manager/routes"
	"github.com/AyoobRukabi/go-task-manager/handlers"
)

func main() {
	r := gin.Default()
	db.ConnectDatabase() // Connect to Postgres/GORM

	// Week 1 routes
	r.GET("/ping", handlers.PingHandler)
	r.GET("/hello", handlers.HelloHandler)

	// Register user routes
    routes.UserRoutes(r) 

	// Week 3 Task routes
	routes.TaskRoutes(r)

	r.Run(":8080")
}
