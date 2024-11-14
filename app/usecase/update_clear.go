package usecase

import (
	"github.com/YKhm20020/Backend-Tacktail/domain"
)

type UpdateClear struct {
	repo domain.CocktailRepository
}

func NewUpdateClear(repo domain.CocktailRepository) UpdateClear {
	return UpdateClear{
		repo: repo,
	}
}
