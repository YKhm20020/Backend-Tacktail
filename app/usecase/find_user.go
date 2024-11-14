package usecase

import (
	"github.com/YKhm20020/Backend-Tacktail/domain"
)

type FindUser struct {
	repo domain.UserRepository
}

func NewFindUser(
	repo domain.UserRepository,
) FindUser {
	return FindUser{
		repo: repo,
	}
}
