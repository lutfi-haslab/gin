package models

type Post struct {
	ID      uint `gorm:"primaryKey"`
	Title   string
	Body    string
	Updated int64 `gorm:"autoUpdateTime"`
	Created int64 `gorm:"autoCreateTime"`
}