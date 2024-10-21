package domain

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint        `json:"user_id" gorm:"not null"`
	Status     string      `json:"status" gorm:"not null"`
	TotalPrice float64     `json:"total_price" gorm:"not null"`
	OrderItems []OrderItem `json:"order_items,omitempty" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	UnitPrice float64 `json:"unit_price" gorm:"not null"`
}

type OrderRepository interface {
	Create(order *Order) error
	GetByID(id uint) (*Order, error)
	GetAll() ([]Order, error)
	GetUserOrders(userID uint) ([]Order, error)
	Update(order *Order) error
	Delete(id uint) error
}
