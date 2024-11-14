package controller

import (
	"errors"
	"net/http"

	"github.com/YKhm20020/Backend-Tacktail/usecase"
	"github.com/gin-gonic/gin"
)

type FindUser struct {
	uc usecase.FindUser
}

func NewFindUser(uc usecase.FindUser) FindUser {
	return FindUser{
		uc: uc,
	}
}

func (con FindUser) Execute(ctx *gin.Context) {
	var input usecase.FindUserInput

	// ヘッダー内のtokenからユーザーIDを取得
	user_id, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("ヘッダー内のトークンからユーザーIDを取得できませんでした")})
		return
	}

	input.ID = user_id.(string)
	output, err := con.uc.Execute(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}
