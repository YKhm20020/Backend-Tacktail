package usecase

import (
	"github.com/KakinokiKanta/go-dev-template/domain"
	"github.com/oklog/ulid/v2"
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
