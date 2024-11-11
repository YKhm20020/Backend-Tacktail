package repository

type dbUser struct {
	id       string
	name     string
	password string
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

type dbImage struct {
	id         string
	cocktailID string
	userID     string
	image      string
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
