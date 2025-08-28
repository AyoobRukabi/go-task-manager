package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/AyoobRukabi/go-task-manager/handlers"
)

func UserRoutes(r *gin.Engine) {
    r.POST("/register", handlers.RegisterUser)
    r.POST("/login", handlers.LoginUser)
}
