package usecase

import (
	"github.com/YKhm20020/Backend-Tacktail/domain"
)

type FindCocktail struct {
	repo domain.CocktailRepository
	// materialRepo domain.MaterialRepository
}

type FindCocktailInput struct {
	ID string
}

type FindCocktailOutput struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
	Materials   []Material `json:"materials"`
}

func NewFindCocktail(
	repo domain.CocktailRepository,
	// materialRepo domain.MaterialRepository,
) FindCocktail {
	return FindCocktail{
		repo: repo,
		// materialRepo: materialRepo,
	}
}

func (uc FindCocktail) Execute(input FindCocktailInput) (FindCocktailOutput, error) {
	cocktail, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return FindCocktailOutput{}, err
	}

	outputMaterials := []Material{}
	for _, material := range cocktail.Materials() {
		outputMaterials = append(outputMaterials, Material{
			material.ID(),
			material.Name(),
			material.Description(),
			material.Amount(),
		})
	}

	return FindCocktailOutput{
		ID:          cocktail.ID(),
		Name:        cocktail.Name(),
		Description: cocktail.Description(),
		Image:       cocktail.Image(),
		Materials:   outputMaterials,
	}, nil
}
