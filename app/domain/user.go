package domain

import "golang.org/x/crypto/bcrypt"

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
	// パスワードをハッシュ化
	hash_pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}
	}

	return User{
		id:       id,
		name:     name,
		password: string(hash_pass),
	}
}

func (user User) ID() string {
	return user.id
}

func (user User) Name() string {
	return user.name
}

func (user User) Password() string {
	return user.password
}
