package controller

import (
	"errors"
	"net/http"

	"github.com/YKhm20020/Backend-Tacktail/usecase"
	"github.com/gin-gonic/gin"
)

type UpdateClear struct {
	uc usecase.UpdateClear
}

func NewUpdateClear(
	uc usecase.UpdateClear,
) UpdateClear {
	return UpdateClear{
		uc: uc,
	}
}

func (con UpdateClear) Execute(ctx *gin.Context) {
	var input usecase.UpdateClearInput

	// ヘッダー内のtokenからユーザーIDを取得
	user_id, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("ヘッダー内のトークンからユーザーIDを取得できませんでした")})
		return
	}

	input.UserID = user_id.(string)
	output, err := con.uc.Execute(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, output)
}
