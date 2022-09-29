package config

import (
	"os"

	// go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {

	godotenv.Load("./.env")
	// USER := os.Getenv("USER")
	// PASS := os.Getenv("PASS")
	// HOST := os.Getenv("HOST")
	// DBNAME := os.Getenv("DBNAME")
	
	// // URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)
	// db, err := gorm.Open("mysql", URL) //Local
	db, err := gorm.Open("mysql", os.Getenv("DSN")) //PlanetScale
	
	if err != nil {
			panic(err.Error())
	}
	
	return db
}