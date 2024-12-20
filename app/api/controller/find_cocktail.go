package controller

import (
	"net/http"

	"github.com/YKhm20020/Backend-Tacktail/usecase"
	"github.com/gin-gonic/gin"
)

type FindCocktail struct {
	uc usecase.FindCocktail
}

func NewFindCocktail(uc usecase.FindCocktail) FindCocktail {
	return FindCocktail{
		uc: uc,
	}
}

func (con FindCocktail) Execute(ctx *gin.Context) {
	var input usecase.FindCocktailInput

	// /cocktails/{id} からカクテルIDを取得する
	input.ID = ctx.Param("id")

	// ヘッダー内のtokenからユーザーIDを取得
	user_id, exists := ctx.Get("user_id")
	if exists {
		input.UserID = user_id.(string)
	} else {
		input.UserID = ""
	}

	output, err := con.uc.Execute(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}
