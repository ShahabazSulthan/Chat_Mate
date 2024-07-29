// /internal/usecases/user/interactor.go
package user

import (
	"chatting/pkg/domain"
	"chatting/pkg/repositories"
)

type UserInteractor struct {
	UserRepository repositories.UserRepository
}

func (u *UserInteractor) CreateUser(user *domain.User) error {
	return u.UserRepository.Create(user)
}

func (u *UserInteractor) GetUser(id string) (*domain.User, error) {
	return u.UserRepository.FindByID(id)
}
