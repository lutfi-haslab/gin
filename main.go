package main

import (
	"gin/config"
	"gin/controllers"
	"gin/models"
	index "gin/test"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

			if c.Request.Method == "OPTIONS" {
					c.AbortWithStatus(204)
					return
			}

			c.Next()
	}
}

func SetupRoutes(db *gorm.DB) *gin.Engine {
    r := gin.Default()
		r.Use(CORSMiddleware())
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })		
// 
    r.GET("/task", controllers.FindTasks)
    r.POST("/tasks", controllers.CreateTask)
		r.GET("/tasks/:id", controllers.FindTask)
		r.PATCH("/tasks/:id", controllers.UpdateTask)
		r.DELETE("tasks/:id", controllers.DeleteTask)

		r.GET("/post", controllers.FindPost)
		r.GET("/post/:id", controllers.FindPostId)
		r.POST("/post", controllers.PostCreate)
		r.PUT("/post/:id", controllers.PostUpdate)
		r.DELETE("/post/:id", controllers.PostDelete)
    return r
}

func main() {
	db := config.SetupDB()
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.Post{})

	r := SetupRoutes(db)
	index.Test()
	r.Run()
}
