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

func (cocktail Cocktail) ID() string {
	return cocktail.id
}

func (cocktail Cocktail) Name() string {
	return cocktail.name
}

func (cocktail Cocktail) Image() string {
	return cocktail.image
}

func (cocktail Cocktail) Description() string {
	return cocktail.description
}

func (cocktail Cocktail) Materials() []Material {
	return cocktail.materials
}
