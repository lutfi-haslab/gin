package models

type Task struct {
	ID         uint `gorm:"primaryKey"`
	AssingedTo string
	Task       string
	Deadline   string
	Updated    int64 `gorm:"autoUpdateTime"`
	Created    int64 `gorm:"autoCreateTime"`
}