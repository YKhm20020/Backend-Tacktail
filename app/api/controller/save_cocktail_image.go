package controller

import (
	"errors"
	"net/http"

	"github.com/YKhm20020/Backend-Tacktail/usecase"
	"github.com/gin-gonic/gin"
)

type SaveCocktailImage struct {
	uc usecase.SaveCocktailImage
}

func NewSaveCocktailImage(
	uc usecase.SaveCocktailImage,
) SaveCocktailImage {
	return SaveCocktailImage{
		uc: uc,
	}
}

func (con SaveCocktailImage) Execute(ctx *gin.Context) {
	var input usecase.SaveCocktailImageInput

	// ヘッダー内のtokenからユーザーIDを取得
	user_id, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("ヘッダー内のトークンからユーザーIDを取得できませんでした")})
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user_id.(string)

	cocktailImage, err := con.uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, cocktailImage)
}
