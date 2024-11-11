package api

import (
	"database/sql"

	"github.com/YKhm20020/Backend-Tacktail/api/controller"
	"github.com/YKhm20020/Backend-Tacktail/infrastructure/repository"
	"github.com/YKhm20020/Backend-Tacktail/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) {
	r := gin.Default()

	// usersに対するリポジトリ
	userRepository := repository.NewUserRepository(db)
	// usersに対するユースケース
	createUserUc := usecase.NewCreateUser(userRepository)
	loginUc := usecase.NewLogin(userRepository)
	// usersに対するコントローラー
	createUserCon := controller.NewCreateUser(createUserUc)
	loginCon := controller.NewLogin(loginUc)

	// /login
	r.POST("/login", loginCon.Execute)

	// /users
	r.POST("/users", createUserCon.Execute)

	r.Run()
}
