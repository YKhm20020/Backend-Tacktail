package usecase

import (
	"github.com/YKhm20020/Backend-Tacktail/domain"
)

type UpdateClear struct {
	repo domain.CocktailRepository
}

type UpdateClearInput struct {
	UserID string
}

type UpdateClearOutput struct {
	UserID string `json:"user_id"`
	Story  int    `json:"story"`
}

func NewUpdateClear(repo domain.CocktailRepository) UpdateClear {
	return UpdateClear{
		repo: repo,
	}
}
