package api

import (
	"database/sql"
	"time"

	"github.com/YKhm20020/Backend-Tacktail/api/controller"
	"github.com/YKhm20020/Backend-Tacktail/api/middleware"
	"github.com/YKhm20020/Backend-Tacktail/infrastructure/repository"
	"github.com/YKhm20020/Backend-Tacktail/usecase"
	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		// 許可するHTTPメソッド一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可するHTTPリクエストヘッダ一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			// "Content-Length",
			// "Accept-Encoding",
			// "Authorization",
			// "accessToken",
			"Set-Cookie",
			"Cookie",
		},
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://frontend-festival-booth.vercel.app",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true, // 本番環境でないと動かない
		MaxAge:           24 * time.Hour,
	}))

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
	r.GET("/cocktails/list", findCocktailListCon.Execute)          // 認証なしでカクテル一覧取得
	authRouter.GET("/cocktails/list", findCocktailListCon.Execute) // 認証ありでカクテル一覧取得

	r.Run()
}
