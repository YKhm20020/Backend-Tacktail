package controller

import "github.com/YKhm20020/Backend-Tacktail/usecase"

type FindCocktail struct {
	uc usecase.FindCocktail
}

func NewFindCocktail(uc usecase.FindCocktail) FindCocktail {
	return FindCocktail{
		uc: uc,
	}
}
