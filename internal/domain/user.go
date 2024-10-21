package domain

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name     string    `json:"name" gorm:"not null"`
    Email    string    `json:"email" gorm:"uniqueIndex;not null"`
    Password string    `json:"-" gorm:"not null"`
    Orders   []Order   `json:"orders,omitempty" gorm:"foreignKey:UserID"`
}

type UserRepository interface {
    Create(user *User) error
    GetByID(id uint) (*User, error)
    GetAll() ([]User, error)
    Update(user *User) error
    Delete(id uint) error
}
