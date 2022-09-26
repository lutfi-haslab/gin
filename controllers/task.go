// controllers/books.go

package controllers

import (
	"gin/config"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /tasks
// Get all tasks
func FindTasks(c *gin.Context) {
	db := config.SetupDB()
	var tasks []models.Task
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// POST /tasks
// Create new task
func CreateTask(c *gin.Context) {

	var body struct {
		AssingedTo string
		Task       string
		Deadline   string
	}

	c.Bind(&body)
	db := config.SetupDB()
	task := models.Task{AssingedTo: body.AssingedTo,Task: body.Task, Deadline: body.Deadline }
	result := db.Create(&task)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// GET /tasks/:id
// Find a task
func FindTask(c *gin.Context) { 

	id := c.Param("id")

	var task models.Task
	db := config.SetupDB()
	db.First(&task, id)
	
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// PATCH /tasks/:id
// Update a task
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		AssingedTo string
		Task       string
		Deadline   string
	}
	c.Bind(&body)
	var task models.Task
	db := config.SetupDB()
	db.First(&task, id)

	db.Model(&task).Updates(models.Task{
		AssingedTo: body.AssingedTo,
		Task: body.Task, 
		Deadline: body.Deadline,
	})
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DELETE /tasks/:id
// Delete a task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	db := config.SetupDB()
	db.Delete(&models.Task{}, id)
	
	c.JSON(200, gin.H{
		"msg": "Data Deleted",
	})
}