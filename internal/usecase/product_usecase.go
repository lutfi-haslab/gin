package usecase

import "clean-api/internal/domain"

type productUseCase struct {
	productRepo domain.ProductRepository
}

func NewProductUseCase(productRepo domain.ProductRepository) domain.ProductRepository {
	return &productUseCase{productRepo}
}

func (u *productUseCase) Create(product *domain.Product) error {
	return u.productRepo.Create(product)
}

func (u *productUseCase) GetByID(id uint) (*domain.Product, error) {
	return u.productRepo.GetByID(id)
}

func (u *productUseCase) GetAll() ([]domain.Product, error) {
	return u.productRepo.GetAll()
}

func (u *productUseCase) Update(product *domain.Product) error {
	return u.productRepo.Update(product)
}

func (u *productUseCase) Delete(id uint) error {
	return u.productRepo.Delete(id)
}
