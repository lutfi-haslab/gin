package main

import (
   "fmt"
    _ "github.com/go-sql-driver/mysql"
    "time"
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// MODELS

type Task struct {
    ID         uint      `json:"id" gorm:"primary_key"`
	AssingedTo string    `json:"assignedTo"`
	Task       string    `json:"task"`
	Deadline   time.Time `json:"deadline"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type CreateTaskInput struct {
    AssingedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline`
}

type UpdateTaskInput struct {
	AssingedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline`
}

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
    USER := "root"
    PASS := "IoA6Ei5fzqwZkYfxXx6a"
    HOST := "containers-us-west-74.railway.app"
    PORT := "7373"
    DBNAME := "railway"
    
    URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
    db, err := gorm.Open("mysql", URL)
    
    if err != nil {
        panic(err.Error())
    }
    
    return db
}

// CONTROLLERrr

// GET /task
// Get all tasks
func FindTasks(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var tasks []Task
    db.Find(&tasks)

    c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// POST /tasks
// Create new task
func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	// Create task
	task := Task{AssingedTo: input.AssingedTo, Task: input.Task, Deadline: deadline}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// GET /tasks/:id
// Find a task
func FindTask(c *gin.Context) { // Get model if exist
	var task Task

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// PATCH /tasks/:id
// Update a task
func UpdateTask(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var task Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	var updatedInput Task
	updatedInput.Deadline = deadline
	updatedInput.AssingedTo = input.AssingedTo
	updatedInput.Task = input.Task

	db.Model(&task).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DELETE /tasks/:id
// Delete a task
func DeleteTask(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book Task
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func SetupRoutes(db *gorm.DB) *gin.Engine {
    r := gin.Default()
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })

    r.GET("/task", FindTasks)
    r.POST("/tasks", CreateTask)
	r.GET("/tasks/:id", FindTask)
	r.PATCH("/tasks/:id", UpdateTask)
	r.DELETE("tasks/:id", DeleteTask)


    return r
}

func main() {
	db := SetupDB()
    db.AutoMigrate(&Task{})

    r := SetupRoutes(db)
	r.Run()
}
