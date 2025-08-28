package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Handles GET /ping
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// Handles GET /hello
func HelloHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	c.JSON(http.StatusOK, gin.H{"message": "Hello " + name})
}
