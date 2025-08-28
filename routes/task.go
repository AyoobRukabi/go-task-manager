package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/AyoobRukabi/go-task-manager/handlers"
	"github.com/AyoobRukabi/go-task-manager/middleware"
)

func TaskRoutes(r *gin.Engine) {
	// Group routes that require authentication
	taskGroup := r.Group("/tasks")
	taskGroup.Use(middleware.AuthMiddleware()) // <-- apply JWT middleware
	{
		taskGroup.POST("/", handlers.CreateTask)
		taskGroup.GET("/", handlers.GetTasks)
		taskGroup.GET("/:id", handlers.GetTask)
		taskGroup.PUT("/:id", handlers.UpdateTask)
		taskGroup.DELETE("/:id", handlers.DeleteTask)
	}
}
