package usecase

import (
	"fmt"

	"github.com/YKhm20020/Backend-Tacktail/domain"
)

type FindCocktailList struct {
	repo domain.CocktailRepository
	// materialRepo domain.MaterialRepository
}

type FindCocktailListInput struct {
	MaterialIDs []string
	UserID      string
}

type FindCocktailListOutput struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
	Materials   []Material `json:"materials"`
}

type Material struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

func NewFindCocktailList(
	repo domain.CocktailRepository,
	// materialRepo domain.MaterialRepository,
) FindCocktailList {
	return FindCocktailList{
		repo: repo,
		// materialRepo: materialRepo,
	}
}

func (uc FindCocktailList) Execute(input FindCocktailListInput) (FindCocktailListOutput, error) {
	var cocktails []domain.Cocktail
	var err error

	if len(input.MaterialIDs) <= 0 {
		fmt.Println("FindAllが呼び出されたよ")
		cocktails, err = uc.repo.FindAll(input.UserID)
	} else {
		fmt.Println("FindByMaterialsが呼び出されたよ")
		cocktails, err = uc.repo.FindByMaterials(input.UserID, input.MaterialIDs)
	}

	if err != nil {
		return nil, err
	}

	return cocktails, nil
}
