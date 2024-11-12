package repository

import (
	"database/sql"
	"fmt"

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

	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return []domain.Cocktail{}, err
	}

	// カクテル一覧を格納する変数
	var cocktails []domain.Cocktail
	// カクテルに対する材料一覧を格納する変数
	var materials []domain.Material

	// クエリの実行結果をバインドするための変数
	var dbCocktail dbCocktail
	var dbMaterial dbMaterial
	var dbCocktailImage dbCocktailImage
	var dbRecipe dbRecipe
	// カクテルに材料のリストを持たせるため、カクテルIDは別で宣言
	cocktailID := ""

	for rows.Next() {
		fmt.Println(cocktailID)
		rows.Scan(
			&dbCocktail.id,
			&dbCocktail.name,
			&dbCocktail.description,
			&dbCocktailImage.image,
			&dbMaterial.id,
			&dbMaterial.name,
			&dbMaterial.description,
			&dbRecipe.amount,
		)

		fmt.Println(
			dbCocktail.id,
			dbCocktail.name,
			dbCocktail.description,
			dbCocktailImage.image,
			dbMaterial.id,
			dbMaterial.name,
			dbMaterial.description,
			dbRecipe.amount,
		)

		if cocktailID == "" || dbCocktail.id == cocktailID {
			material := domain.NewMaterial(dbMaterial.id, dbMaterial.name, dbMaterial.description, dbRecipe.amount)
			materials = append(materials, material)
			fmt.Println(material)
		} else {
			cocktail := domain.NewCocktail(dbCocktail.id, dbCocktail.name, dbCocktail.description, dbCocktailImage.image, materials)
			cocktails = append(cocktails, cocktail)

			cocktailID = dbCocktail.id
			materials = nil

			fmt.Println(cocktail)
		}
	}

	// TODO: 確認用、後で消す
	for _, contents := range cocktails {
		fmt.Println(contents)
	}

	return cocktails, nil
}

func (repo CocktailRepository) FindByMaterials(
	userID string,
	materialIDs []string,
) ([]domain.Cocktail, error) {
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

	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return []domain.Cocktail{}, err
	}

	// カクテル一覧を格納する変数
	var cocktails []domain.Cocktail
	// カクテルに対する材料一覧を格納する変数
	var materials []domain.Material

	// クエリの実行結果をバインドするための変数
	var dbCocktail dbCocktail
	var dbMaterial dbMaterial
	var dbCocktailImage dbCocktailImage
	var dbRecipe dbRecipe
	// カクテルに材料のリストを持たせるため、カクテルIDは別で宣言
	cocktailID := ""

	for rows.Next() {
		rows.Scan(
			&dbCocktail.id,
			&dbCocktail.name,
			&dbCocktail.description,
			&dbCocktailImage.image,
			&dbMaterial.id,
			&dbMaterial.name,
			&dbMaterial.description,
			&dbRecipe.amount,
		)

		if cocktailID == "" || dbCocktail.id == cocktailID {
			material := domain.NewMaterial(dbMaterial.id, dbMaterial.name, dbMaterial.description, dbRecipe.amount)
			materials = append(materials, material)
		} else {
			cocktail := domain.NewCocktail(dbCocktail.id, dbCocktail.name, dbCocktail.description, dbCocktailImage.image, materials)
			cocktails = append(cocktails, cocktail)

			cocktailID = dbCocktail.id
			materials = nil
		}
	}

	// TODO: 確認用、後で消す
	for _, contents := range cocktails {
		fmt.Println(contents)
	}

	return cocktails, nil
}
