package controller

import (
	"net/http"

	"github.com/YKhm20020/Backend-Tacktail/usecase"
	"github.com/gin-gonic/gin"
)

type FindCocktailList struct {
	uc usecase.FindCocktailList
}

func NewFindCocktailList(uc usecase.FindCocktailList) FindCocktailList {
	return FindCocktailList{
		uc: uc,
	}
}

func (con FindCocktailList) Execute(ctx *gin.Context) {
	var input usecase.FindCocktailListInput

	// /cocktails/list?materials={material_id}?materials={material_id} でフィルターできるようにするため
	input.MaterialIDs = ctx.QueryArray("materials")

	// cookieからユーザーIDを取得
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
