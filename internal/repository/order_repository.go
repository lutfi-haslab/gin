package repository

import (
	"clean-api/internal/domain"

	"gorm.io/gorm"
)

type gormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &gormOrderRepository{db}
}

func (r *gormOrderRepository) Create(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *gormOrderRepository) GetByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("OrderItems").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *gormOrderRepository) GetAll() ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Preload("OrderItems").Find(&orders).Error
	return orders, err
}

func (r *gormOrderRepository) GetUserOrders(userID uint) ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Preload("OrderItems").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (r *gormOrderRepository) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}

func (r *gormOrderRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Order{}, id).Error
}
