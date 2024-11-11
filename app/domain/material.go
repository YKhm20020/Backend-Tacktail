package domain

type Material struct {
	id          string
	name        string
	description string
	amount      int
}

type MaterialRepository interface {
	Find(string) (Material, error)      // 材料IDを指定
	Finds([]string) ([]Material, error) // 複数の材料IDを指定
}

func NewMaterial(
	id string,
	name string,
	description string,
) Material {
	return Material{
		id:          id,
		name:        name,
		description: description,
	}
}

func (material Material) ID() string {
	return material.id
}

func (material Material) Name() string {
	return material.name
}

func (material Material) Description() string {
	return material.description
}

func (material Material) Amount() int {
	return material.amount
}
