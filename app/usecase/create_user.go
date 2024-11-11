package usecase

import (
	"errors"

	"github.com/YKhm20020/Backend-Tacktail/domain"
)

type CreateUser struct {
	repo domain.UserRepository
}

func NewCreateUser(repo domain.UserRepository) CreateUser {
	return CreateUser{
		repo: repo,
	}
}

type CreateUserInput struct {
	Name     string `json:"name" binding:"required,min=1,max=20"`
	Password string `json:"password" binding:"required,min=1"`
}

type CreateUserOutput struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (uc CreateUser) Execute(input CreateUserInput) (CreateUserOutput, error) {
	_, err := uc.repo.FindByName(input.Name)

	if err == nil {
		return CreateUserOutput{}, errors.New("this name is already registered")
	}

	if err.Error() != "not found user" {
		return CreateUserOutput{}, err
	}

	user := domain.NewUser(input.Name, input.Password)

	createdUser, err := uc.repo.Create(user)

	if err != nil {
		return CreateUserOutput{}, err
	}

	return CreateUserOutput{
		createdUser.ID(),
		createdUser.Name(),
		createdUser.Password(),
	}, nil
}
