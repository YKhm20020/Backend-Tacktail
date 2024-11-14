package usecase

import "github.com/YKhm20020/Backend-Tacktail/domain"

type SaveCocktailImage struct {
	repo domain.CocktailRepository
}

func NewSaveCocktailImage(repo domain.CocktailRepository) SaveCocktailImage {
	return SaveCocktailImage{
		repo: repo,
	}
}
