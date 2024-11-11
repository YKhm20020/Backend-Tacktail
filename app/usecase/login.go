package usecase

import (
	"github.com/YKhm20020/Backend-Tacktail/domain"
)

type Login struct {
	repo domain.UserRepository
}

type LoginInput struct {
	Name     string `json:"name" binding:"required,min=1,max=20"`
	Password string `json:"password" binding:"required,min=1"`
}

func NewLogin(repo domain.UserRepository) Login {
	return Login{
		repo: repo,
	}
}
