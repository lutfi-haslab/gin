package usecase

import (
	"clean-api/internal/domain"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserRepository {
	return &userUseCase{userRepo}
}

func (u *userUseCase) Create(user *domain.User) error {
	return u.userRepo.Create(user)
}

func (u *userUseCase) GetByID(id uint) (*domain.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUseCase) GetAll() ([]domain.User, error) {
	return u.userRepo.GetAll()
}

func (u *userUseCase) Update(user *domain.User) error {
	return u.userRepo.Update(user)
}

func (u *userUseCase) Delete(id uint) error {
	return u.userRepo.Delete(id)
}
