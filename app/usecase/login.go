package usecase

import (
	"github.com/YKhm20020/Backend-Tacktail/domain"
)

type Login struct {
	repo domain.UserRepository
}

func NewLogin(repo domain.UserRepository) Login {
	return Login{
		repo: repo,
	}
}
