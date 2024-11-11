package repository

import (
	"database/sql"

	"github.com/YKhm20020/Backend-Tacktail/domain"
)

type CocktailRepository struct {
	db *sql.DB
}

func NewCocktailRepository(db *sql.DB) domain.CocktailRepository {
	return CocktailRepository{
		db: db,
	}
}

func (repo CocktailRepository) UpdateCocktailImage(
	userID string,
	cocktailID string,
	image string,
) error {
	return nil
}

func (repo CocktailRepository) FindAll(userID string) ([]domain.Cocktail, error) {
	return []domain.Cocktail{}, nil
}

func (repo CocktailRepository) FindByMaterials(
	userID string,
	materialIDs []string,
) ([]domain.Cocktail, error) {
	return []domain.Cocktail{}, nil
}
