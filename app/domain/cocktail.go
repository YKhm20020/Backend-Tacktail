package domain

type Cocktail struct {
	id          string
	name        string
	image       string // ユーザー固有の値
	description string
	materials   []Material
}

type CocktailRepository interface {
	UpdateCocktailImage(string, string, string) error     // 引数はユーザーID、カクテルID、画像パス
	FindAll(string) ([]Cocktail, error)                   // 引数はユーザーIDで、すべてのカクテルを取得
	FindByMaterials(string, []string) ([]Cocktail, error) // 引数はユーザーIDと材料IDのリスト
	// FindByID(string) (Cocktail, error)                    // 引数はカクテルID
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
