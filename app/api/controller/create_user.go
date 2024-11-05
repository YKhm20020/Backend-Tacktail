package controller

import (
	"net/http"

	"github.com/KakinokiKanta/go-dev-template/usecase"
	"github.com/gin-gonic/gin"
)

type CreateUser struct {
	uc usecase.CreateUser
}

func NewCreateUser(uc usecase.CreateUser) CreateUser {
	return CreateUser{
		uc: uc,
	}
}
