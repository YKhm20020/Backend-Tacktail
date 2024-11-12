package api

import (
	"database/sql"

	"github.com/YKhm20020/Backend-Tacktail/api/controller"
	"github.com/YKhm20020/Backend-Tacktail/api/middleware"
	"github.com/YKhm20020/Backend-Tacktail/infrastructure/repository"
	"github.com/YKhm20020/Backend-Tacktail/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) {
	// users
	// リポジトリ
	userRepository := repository.NewUserRepository(db)
	// ユースケース
	createUserUc := usecase.NewCreateUser(userRepository)
	loginUc := usecase.NewLogin(userRepository)
	// コントローラ
	createUserCon := controller.NewCreateUser(createUserUc)
	loginCon := controller.NewLogin(loginUc)

	r := gin.Default()

	// cocktails
	// リポジトリ
	cocktailRepository := repository.NewCocktailRepository(db)
	// ユースケース
	findCocktailListUc := usecase.NewFindCocktailList(cocktailRepository)
	// コントローラ
	findCocktailListCon := controller.NewFindCocktailList(findCocktailListUc)

	// 認証を必要とするルーティングを定義
	authRouter := r.Group("/auth", middleware.AuthJWT())

	// /login
	r.POST("/login", loginCon.Execute)

	// /users
	r.POST("/users", createUserCon.Execute)

	// /cocktails
	r.GET("/cocktails", findCocktailListCon.Execute)          // 認証なしでカクテル一覧取得
	authRouter.GET("/cocktails", findCocktailListCon.Execute) // 認証ありでカクテル一覧取得

	r.Run()
}
