package api

import (
	"database/sql"

	"github.com/KakinokiKanta/go-dev-template/api/controller"
	"github.com/KakinokiKanta/go-dev-template/infrastructure/repository"
	"github.com/KakinokiKanta/go-dev-template/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) {
	r := gin.Default()

	userRepository := repository.NewUserRepository(db)
	createUserUc := usecase.NewCreateUser(userRepository)
	createUserCon := controller.NewCreateUser(createUserUc)

	r.POST("/users", createUserCon.Execute)
	r.Run()
}
