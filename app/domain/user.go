package domain

import (
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id       string
	name     string
	password string
	story    int
}

type UserRepository interface {
	Create(User) (User, error)
	FindByName(string) (User, error)
}

func NewUser(
	name string,
	password string,
) User {
	// IDを生成
	id := ulid.Make().String()

	// パスワードをハッシュ化
	hash_pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}
	}

	// ストーリーモードのクリア状況を初期化
	story := 0

	return newUser(id, name, string(hash_pass), story)
}

func ReUser(
	id string,
	name string,
	password string,
	story int,
) User {
	return newUser(id, name, password, story)
}

func newUser(
	id string,
	name string,
	password string,
	story int,
) User {
	return User{
		id:       id,
		name:     name,
		password: password,
		story:    story,
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

func (user User) Story() int {
	return user.story
}

func (user User) IsValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password))

	return err != bcrypt.ErrMismatchedHashAndPassword
}
