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
	query := `
		SELECT
		cocktails.id, cocktails.name, cocktails.description, cocktail_images.image,
			materials.id, materials.name, materials.description, recipes.amount
		FROM
		cocktails
		LEFT JOIN
		cocktail_images
		ON
		cocktails.id = cocktail_images.cocktailID
		AND
		cocktail_images.userID = $1
		INNER JOIN
		recipes
		ON
		cocktails.id = recipes.cocktailID
		INNER JOIN
		materials
		ON
		recipes.materialID = materials.id
	`

	return []domain.Cocktail{}, nil
}

func (repo CocktailRepository) FindByMaterials(
	userID string,
	materialIDs []string,
) ([]domain.Cocktail, error) {
	return []domain.Cocktail{}, nil
}
