package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/AyoobRukabi/go-task-manager/handlers"
)

func TaskRoutes(r *gin.Engine) {
	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetTasks)
	r.GET("/tasks/:id", handlers.GetTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)
}
