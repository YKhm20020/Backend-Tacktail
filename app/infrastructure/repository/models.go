package repository

import "database/sql"

type dbUser struct {
	id       string
	name     string
	password string
	story    int
}

type dbCocktail struct {
	id          string
	name        string
	description string
}

type dbRecipe struct {
	id         string
	cocktailID string
	materialID string
	amount     int
}

type dbCocktailImage struct {
	id         string
	cocktailID string
	userID     string
	image      sql.NullString
}

type dbMaterial struct {
	id          string
	name        string
	description string
}

type dbStory struct {
	id         string
	cocktailID string
	trivia     string
	day        int
}
