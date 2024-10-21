package repository

import (
	"clean-api/internal/domain"

	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) domain.UserRepository {
	return &gormUserRepository{db}
}

func (r *gormUserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *gormUserRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Orders.OrderItems").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *gormUserRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *gormUserRepository) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}
