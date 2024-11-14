package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/YKhm20020/Backend-Tacktail/domain"
	"github.com/lib/pq"
)

type CocktailRepository struct {
	db *sql.DB
}

func NewCocktailRepository(db *sql.DB) domain.CocktailRepository {
	return CocktailRepository{
		db: db,
	}
}

func (repo CocktailRepository) InsertCocktailImage(
	cocktailImageID string,
	userID string,
	cocktailID string,
	image string,
) (string, error) {
	query := `
		INSERT INTO cocktail_images (id, cocktailID, userID, image) VALUES ($1, $2, $3, $4);
	`
	_, err := repo.db.Exec(query, cocktailImageID, userID, cocktailID, image)

	if err != nil {
		return "", err
	}

	return cocktailImageID, nil
}

func (repo CocktailRepository) UpdateCocktailImage(
	cocktailImageID string,
	image string,
) (string, error) {
	query := `
		UPDATE cocktail_images SET image = $1 WHERE id = $2;
	`
	_, err := repo.db.Exec(query, image, cocktailImageID)

	if err != nil {
		return "", err
	}

	return cocktailImageID, nil
}

func (repo CocktailRepository) FindAll(userID string) (map[string]domain.Cocktail, error) {
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
		// return []domain.Cocktail{}, err
		return nil, err
	}

	// クエリの実行結果をバインドするための変数
	var dbCocktail dbCocktail
	var dbMaterial dbMaterial
	var dbCocktailImage dbCocktailImage
	var dbRecipe dbRecipe

	// カクテルIDからカクテル構造体へのマップ
	cocktailMap := make(map[string]domain.Cocktail)
	// カクテルIDから材料構造体へのマップ
	materialMap := make(map[string][]domain.Material)

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

		// 材料ドメインのオブジェクトを生成
		material := domain.NewMaterial(dbMaterial.id, dbMaterial.name, dbMaterial.description, dbRecipe.amount)
		// マップへ紐づけ
		materialMap[dbCocktail.id] = append(materialMap[dbCocktail.id], material)

		// カクテルドメインのオブジェクトを生成
		cocktail := domain.NewCocktail(dbCocktail.id, dbCocktail.name, dbCocktail.description, dbCocktailImage.image.String, materialMap[dbCocktail.id])
		// マップへ紐づけ
		cocktailMap[dbCocktail.id] = cocktail
	}

	return cocktailMap, nil
}

func (repo CocktailRepository) FindByMaterials(
	userID string,
	materialIDs []string,
) (map[string]domain.Cocktail, error) {
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
		WHERE
			cocktails.id 
		IN
			(
				SELECT
					recipes.cocktailID
				FROM
					recipes
				WHERE
					recipes.materialID = ANY($2)
			);
	`

	fmt.Println(materialIDs)

	rows, err := repo.db.Query(query, userID, pq.Array(materialIDs))
	if err != nil {
		// return []domain.Cocktail{}, err
		return nil, err
	}

	// クエリの実行結果をバインドするための変数
	var dbCocktail dbCocktail
	var dbMaterial dbMaterial
	var dbCocktailImage dbCocktailImage
	var dbRecipe dbRecipe

	// カクテルIDからカクテル構造体へのマップ
	cocktailMap := make(map[string]domain.Cocktail)
	// カクテルIDから材料構造体へのマップ
	materialMap := make(map[string][]domain.Material)

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

		// 材料ドメインのオブジェクトを生成
		material := domain.NewMaterial(dbMaterial.id, dbMaterial.name, dbMaterial.description, dbRecipe.amount)
		// マップへ紐づけ
		materialMap[dbCocktail.id] = append(materialMap[dbCocktail.id], material)

		// カクテルドメインのオブジェクトを生成
		cocktail := domain.NewCocktail(dbCocktail.id, dbCocktail.name, dbCocktail.description, dbCocktailImage.image.String, materialMap[dbCocktail.id])
		// マップへ紐づけ
		cocktailMap[dbCocktail.id] = cocktail
	}

	return cocktailMap, nil
}

func (repo CocktailRepository) FindImage(
	userID string,
	cocktailID string,
) (string, error) {
	query := `
		SELECT id FROM cocktail_images WHERE cocktailID=$1 AND userID=$2;
	`

	var dbCocktailImage dbCocktailImage
	err := repo.db.QueryRow(query, cocktailID, userID).Scan(&dbCocktailImage.id)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("not found image")
		}
		return "", err
	}

	return cocktailID, nil
}
