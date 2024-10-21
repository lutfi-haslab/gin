package domain

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string      `json:"name" gorm:"not null"`
	Description string      `json:"description"`
	Price       float64     `json:"price" gorm:"not null"`
	Stock       int         `json:"stock" gorm:"not null"`
	OrderItems  []OrderItem `json:"order_items,omitempty" gorm:"foreignKey:ProductID"`
}

type ProductRepository interface {
	Create(product *Product) error
	GetByID(id uint) (*Product, error)
	GetAll() ([]Product, error)
	Update(product *Product) error
	Delete(id uint) error
}
