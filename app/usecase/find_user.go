package usecase

import (
	"github.com/YKhm20020/Backend-Tacktail/domain"
)

type FindUser struct {
	repo domain.UserRepository
}

type FindUserInput struct {
	ID string
}

type FindUserOutput struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Story    int    `json:"story"`
}

func NewFindUser(
	repo domain.UserRepository,
) FindUser {
	return FindUser{
		repo: repo,
	}
}

func (uc FindUser) Execute(input FindUserInput) (FindUserOutput, error) {
	user, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return FindUserOutput{}, err
	}

	return FindUserOutput{
		user.ID(),
		user.Name(),
		user.Password(),
		user.Story(),
	}, nil
}
