package config

import (
	"fmt"

	// go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

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