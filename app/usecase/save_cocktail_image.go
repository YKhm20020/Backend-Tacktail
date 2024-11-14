package usecase

import "github.com/YKhm20020/Backend-Tacktail/domain"

type SaveCocktailImage struct {
	repo domain.CocktailRepository
}

type SaveCocktailImageInput struct {
	CocktailID string `json:"cocktail_id" binding:"required"`
	UserID     string `json:"user_id" binding:"required"`
	Image      string `json:"image" binding:"required"`
}

type SaveCocktailImageOutput struct {
	ImageID    string `json:"image_id"`
	CocktailID string `json:"cocktail_id"`
	UserID     string `json:"user_id"`
	Image      string `json:"image"`
}

func NewSaveCocktailImage(repo domain.CocktailRepository) SaveCocktailImage {
	return SaveCocktailImage{
		repo: repo,
	}
}

// func (uc SaveCocktailImage) Execute()
