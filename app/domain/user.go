package domain

type User struct {
	id       string
	name     string
	password string
}

type UserRepository interface {
	Create(User) (User, error)
	FindByName(string) (User, error)
}

func NewUser(id string, name string, password string) User {
	return User{
		id:       id,
		name:     name,
		password: password,
	}
}
