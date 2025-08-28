package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create a Gin router
	r := gin.Default()

	// Define a simple route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})





	r.GET("/hello", func(c *gin.Context) {
    name := c.DefaultQuery("name", "Guest")
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello " + name,
    })
})
	

// Start server on port 8080
r.Run(":8080")


}

