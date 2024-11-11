package domain

type Cocktail struct {
	id          string
	name        string
	image       string
	description string
	materials   []Material
}

func NewCocktail(
	id string,
	name string,
	image string,
	description string,
	materials []Material,
) Cocktail {
	return Cocktail{
		id:          id,
		name:        name,
		image:       image,
		description: description,
		materials:   materials,
	}
}
