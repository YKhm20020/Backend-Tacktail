package usecase

import (
	"github.com/YKhm20020/Backend-Tacktail/domain"
)

type UpdateClear struct {
	repo domain.UserRepository
}

type UpdateClearInput struct {
	UserID string
}

type UpdateClearOutput struct {
	UserID string `json:"user_id"`
	Story  int    `json:"story"`
}

func NewUpdateClear(repo domain.UserRepository) UpdateClear {
	return UpdateClear{
		repo: repo,
	}
}

func (uc UpdateClear) Execute(input UpdateClearInput) (UpdateClearOutput, error) {
	err := uc.repo.UpdateStory(input.UserID)
	if err != nil {
		return UpdateClearOutput{}, err
	}

	user, err := uc.repo.FindByID(input.UserID)
	if err != nil {
		return UpdateClearOutput{}, err
	}

	return UpdateClearOutput{
		user.ID(),
		user.Story(),
	}, nil
}
