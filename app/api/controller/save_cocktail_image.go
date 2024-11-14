package controller

import "github.com/YKhm20020/Backend-Tacktail/usecase"

type SaveCocktailImage struct {
	uc usecase.SaveCocktailImage
}

func NewSaveCocktailImage(
	uc usecase.SaveCocktailImage,
) SaveCocktailImage {
	return SaveCocktailImage{
		uc: uc,
	}
}
