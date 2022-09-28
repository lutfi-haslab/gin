package main

import (
	"gin/config"
	"gin/controllers"
	"gin/models"
	index "gin/test"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
    r := gin.Default()
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })

		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"https://foo.com"},
			AllowMethods:     []string{"PUT", "PATCH"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
			 return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,
		 }))

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
