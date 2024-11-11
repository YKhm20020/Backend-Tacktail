package domain

type Material struct {
	id          string
	name        string
	description string
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
