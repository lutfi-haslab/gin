package controllers

import (
	"gin/config"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindPost(c *gin.Context){
	var posts []models.Post
	db := config.SetupDB()
	db.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})

}

func FindPostId(c *gin.Context){
	id := c.Param("id")

	var post models.Post
	db := config.SetupDB()
	db.First(&post, id)
	
	c.JSON(http.StatusOK, gin.H{"data": post})
}


func PostCreate(c *gin.Context){
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)
	db := config.SetupDB()
	post := models.Post{Title: body.Title, Body: body.Body}
	result := db.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context){
	id := c.Param("id")

	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)
	var post models.Post
	db := config.SetupDB()
	db.First(&post, id)
	db.Model(&post).Updates(models.Post{
		Title: body.Title, 
		Body: body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context){
	id := c.Param("id")

	db := config.SetupDB()
	db.Delete(&models.Post{}, id)
	
	c.JSON(200, gin.H{
		"msg": "Data Deleted",
	})

}
