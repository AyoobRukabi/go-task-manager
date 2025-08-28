package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/AyoobRukabi/go-task-manager/db"
	"github.com/AyoobRukabi/go-task-manager/models"
)

// Create a new task
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Create(&task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// Get all tasks
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	result := db.DB.Find(&tasks)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// Get single task by ID
func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	result := db.DB.First(&task, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Update task by ID
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := db.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	task.Title = input.Title
	task.Completed = input.Completed
	task.DueDate = input.DueDate

	db.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// Delete task by ID
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := db.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	db.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
