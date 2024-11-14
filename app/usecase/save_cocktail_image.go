package usecase

import (
	"github.com/YKhm20020/Backend-Tacktail/domain"
	"github.com/oklog/ulid/v2"
)

type SaveCocktailImage struct {
	repo domain.CocktailRepository
}

type SaveCocktailImageInput struct {
	UserID     string `json:"user_id" binding:"required"`
	CocktailID string `json:"cocktail_id" binding:"required"`
	Image      string `json:"image" binding:"required"`
}

type SaveCocktailImageOutput struct {
	ImageID    string `json:"image_id"`
	UserID     string `json:"user_id"`
	CocktailID string `json:"cocktail_id"`
	Image      string `json:"image"`
}

func NewSaveCocktailImage(repo domain.CocktailRepository) SaveCocktailImage {
	return SaveCocktailImage{
		repo: repo,
	}
}

func (uc SaveCocktailImage) Execute(input SaveCocktailImageInput) (SaveCocktailImageOutput, error) {
	// カクテルIDとユーザーIDから画像パスがDB内に存在するかチェック
	cocktailImageID, err := uc.repo.FindImage(input.UserID, input.CocktailID)
	if err == nil {
		_, err = uc.repo.UpdateCocktailImage(cocktailImageID, input.Image)
		if err != nil {
			return SaveCocktailImageOutput{}, err
		}

		return SaveCocktailImageOutput{
			cocktailImageID,
			input.UserID,
			input.CocktailID,
			input.Image,
		}, nil
	}

	if err.Error() != "not found image" {
		return SaveCocktailImageOutput{}, err
	}

	cocktailImageID, err = uc.repo.InsertCocktailImage(ulid.Make().String(), input.Image)
	if err != nil {
		return SaveCocktailImageOutput{}, err
	}

	return SaveCocktailImageOutput{
		cocktailImageID,
		input.UserID,
		input.CocktailID,
		input.Image,
	}, nil
}
