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
