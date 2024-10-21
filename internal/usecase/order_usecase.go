package usecase

import "clean-api/internal/domain"

type orderUseCase struct {
	orderRepo domain.OrderRepository
}

func NewOrderUseCase(orderRepo domain.OrderRepository) domain.OrderRepository {
	return &orderUseCase{orderRepo}
}

func (u *orderUseCase) Create(order *domain.Order) error {
	return u.orderRepo.Create(order)
}

func (u *orderUseCase) GetByID(id uint) (*domain.Order, error) {
	return u.orderRepo.GetByID(id)
}

func (u *orderUseCase) GetAll() ([]domain.Order, error) {
	return u.orderRepo.GetAll()
}

func (u *orderUseCase) GetUserOrders(userID uint) ([]domain.Order, error) {
	return u.orderRepo.GetUserOrders(userID)
}

func (u *orderUseCase) Update(order *domain.Order) error {
	return u.orderRepo.Update(order)
}

func (u *orderUseCase) Delete(id uint) error {
	return u.orderRepo.Delete(id)
}
