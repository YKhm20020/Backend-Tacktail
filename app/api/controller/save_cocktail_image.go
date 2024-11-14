package controller

import (
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

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cocktailImage, err := con.uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, cocktailImage)
}
