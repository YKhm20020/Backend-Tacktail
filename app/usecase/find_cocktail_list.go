package usecase

import (
	"fmt"

	"github.com/YKhm20020/Backend-Tacktail/domain"
)

type FindCocktailList struct {
	cocktailRepo domain.CocktailRepository
	materialRepo domain.MaterialRepository
}

type FindCocktailListInput struct {
	MaterialIDs []string
	UserID      string
}

type FindCocktailListOutput []domain.Cocktail

func NewFindCocktailList(
	cocktailRepo domain.CocktailRepository,
	materialRepo domain.MaterialRepository,
) FindCocktailList {
	return FindCocktailList{
		cocktailRepo: cocktailRepo,
		materialRepo: materialRepo,
	}
}

func (uc FindCocktailList) Execute(input FindCocktailListInput) (FindCocktailListOutput, error) {
	var cocktails []domain.Cocktail
	var err error

	if len(input.MaterialIDs) <= 0 {
		fmt.Println("FindAllが呼び出されたよ")
		cocktails, err = uc.cocktailRepo.FindAll(input.UserID)
	} else {
		fmt.Println("FindByMaterialsが呼び出されたよ")
		cocktails, err = uc.cocktailRepo.FindByMaterials(input.UserID, input.MaterialIDs)
	}

	if err != nil {
		return nil, err
	}

	return cocktails, nil
}
