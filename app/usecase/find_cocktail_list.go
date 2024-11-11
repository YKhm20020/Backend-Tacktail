package usecase

import "github.com/YKhm20020/Backend-Tacktail/domain"

type FindCocktailList struct {
	repo domain.CocktailRepository
}

type FindCocktailListInput struct {
	MaterialID []string `json:"material_id"`
	UserID     string   `json:"user_id"`
}

type FindCocktailListOutput []domain.Cocktail

func NewFindCocktailList(repo domain.CocktailRepository) FindCocktailList {
	return FindCocktailList{
		repo: repo,
	}
}
