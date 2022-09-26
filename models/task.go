package models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	AssingedTo string
	Task       string
	Deadline   string
}