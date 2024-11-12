package controller

import "github.com/YKhm20020/Backend-Tacktail/usecase"

type FindCocktailList struct {
	uc usecase.FindCocktailList
}

func NewFindCocktailList(uc usecase.FindCocktailList) FindCocktailList {
	return FindCocktailList{
		uc: uc,
	}
}
