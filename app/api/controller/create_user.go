package controller

import (
	"net/http"

	"github.com/YKhm20020/Backend-Tacktail/usecase"
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

// PostTodo ...
// @Summary 新規Todoを作成
// @Tags Todo
// @Accept  json
// @Produce  json
// @Param title body string true "title"
// @Param body body string true "body"
// @Success 201 {object} responses.SuccessResponse{data=models.Todo}
// @Failure 400 {object} responses.ErrorResponse
// @Router /todos [post]
func (con CreateUser) Execute(ctx *gin.Context) {
	var input usecase.CreateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := con.uc.Execute(input)

	if err != nil {
		if err.Error() == "this name is already registered" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "This name is already registered"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
