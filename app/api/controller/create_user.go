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

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
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
