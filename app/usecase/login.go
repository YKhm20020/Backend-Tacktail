package usecase

import (
	"errors"

	"github.com/YKhm20020/Backend-Tacktail/domain"
)

type Login struct {
	repo domain.UserRepository
}

type LoginInput struct {
	Name     string `json:"name" binding:"required,min=1,max=20"`
	Password string `json:"password" binding:"required,min=1"`
}

type LoginOutput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewLogin(repo domain.UserRepository) Login {
	return Login{
		repo: repo,
	}
}

/*
 * ここではトークンの生成は行わず、ログインの成否だけを判定する
 * これにより、認証手段を変更しやすくしている
 */
func (uc Login) Execute(input LoginInput) (LoginOutput, error) {
	user, err := uc.repo.FindByName(input.Name)
	if err != nil {
		return LoginOutput{}, err
	}

	if !user.IsValidPassword(input.Password) {
		return LoginOutput{}, errors.New("password is incorrect")
	}

	return LoginOutput{
		Id:   user.ID(),
		Name: user.Name(),
	}, nil
}
