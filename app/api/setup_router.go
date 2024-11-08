package api

import (
	"database/sql"

	"github.com/YKhm20020/Backend-Tacktail/api/controller"
	"github.com/YKhm20020/Backend-Tacktail/infrastructure/repository"
	"github.com/YKhm20020/Backend-Tacktail/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	userRepository := repository.NewUserRepository(db)
	createUserUc := usecase.NewCreateUser(userRepository)
	createUserCon := controller.NewCreateUser(createUserUc)

	r.POST("/users", createUserCon.Execute)

	return r
}
