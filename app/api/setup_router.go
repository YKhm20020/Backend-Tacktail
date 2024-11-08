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

	// swagger用エンドポイント
	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	userRepository := repository.NewUserRepository(db)
	createUserUc := usecase.NewCreateUser(userRepository)
	createUserCon := controller.NewCreateUser(createUserUc)

	r.POST("/users", createUserCon.Execute)
	r.Run()
}
