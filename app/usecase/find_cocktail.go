package usecase

import "github.com/YKhm20020/Backend-Tacktail/domain"

type FindCocktail struct {
	repo domain.CocktailRepository
}

func NewFindCocktail(repo domain.CocktailRepository) FindCocktail {
	return FindCocktail{
		repo: repo,
	}
}
