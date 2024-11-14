package controller

import (
	"github.com/YKhm20020/Backend-Tacktail/usecase"
)

type FindUser struct {
	uc usecase.FindUser
}

func NewFindUser(uc usecase.FindUser) FindUser {
	return FindUser{
		uc: uc,
	}
}
