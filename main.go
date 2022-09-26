package main

import (
	"gin/config"
	"gin/controllers"
	"gin/models"
	index "gin/test"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Test Upload File
// Test upload file
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")

	// The file cannot be received.
	if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "No file is received",
			})
			return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, "/aset" + newFileName); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to save the file",
			})
			return
	}

	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
			"message": "Your file has been successfully uploaded.",
	})
}

func SetupRoutes(db *gorm.DB) *gin.Engine {
    r := gin.Default()
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })

    r.GET("/task", controllers.FindTasks)
    r.POST("/tasks", controllers.CreateTask)
		r.GET("/tasks/:id", controllers.FindTask)
		r.PATCH("/tasks/:id", controllers.UpdateTask)
		r.DELETE("tasks/:id", controllers.DeleteTask)
		r.POST("/upload", Upload)


    return r
}

func main() {
	db := config.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := SetupRoutes(db)
	index.Test()
	r.Run()
}
