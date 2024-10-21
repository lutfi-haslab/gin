package repository

import (
	"clean-api/internal/domain"

	"gorm.io/gorm"
)

type gormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) domain.ProductRepository {
	return &gormProductRepository{db}
}

func (r *gormProductRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *gormProductRepository) GetByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *gormProductRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *gormProductRepository) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}

func (r *gormProductRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Product{}, id).Error
}
