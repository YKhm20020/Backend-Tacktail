package controller

import "github.com/YKhm20020/Backend-Tacktail/usecase"

type Login struct {
	uc usecase.Login
}

func NewLogin(uc usecase.Login) Login {
	return Login{
		uc: uc,
	}
}
