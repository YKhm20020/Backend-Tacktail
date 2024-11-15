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
	updateClearUc := usecase.NewUpdateClear(userRepository)
	loginUc := usecase.NewLogin(userRepository)
	findUserUc := usecase.NewFindUser(userRepository)
	// コントローラ
	createUserCon := controller.NewCreateUser(createUserUc)
	updateClearCon := controller.NewUpdateClear(updateClearUc)
	loginCon := controller.NewLogin(loginUc)
	findUserCon := controller.NewFindUser(findUserUc)

	// cocktails
	// リポジトリ
	cocktailRepository := repository.NewCocktailRepository(db)
	// ユースケース
	findCocktailUc := usecase.NewFindCocktail(cocktailRepository)
	findCocktailListUc := usecase.NewFindCocktailList(cocktailRepository)
	saveCocktailImageUd := usecase.NewSaveCocktailImage(cocktailRepository)
	// コントローラ
	findCocktailCon := controller.NewFindCocktail(findCocktailUc)
	findCocktailListCon := controller.NewFindCocktailList(findCocktailListUc)
	saveCocktailImageCon := controller.NewSaveCocktailImage(saveCocktailImageUd)

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
			"Authorization",
			// "accessToken",
			// "Set-Cookie",
			// "Cookie",
		},
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://frontend-tacktail.vercel.app/",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true, // 本番環境でないと動かない
		MaxAge:           24 * time.Hour,
	}))

	// 認証を必要とするルーティングを定義
	authRouter := r.Group("/auth", middleware.AuthJWT())

	// /login
	r.POST("/login", loginCon.Execute)

	// /users
	r.POST("/users", createUserCon.Execute)
	authRouter.GET("/users", findUserCon.Execute)
	authRouter.POST("/users/story", updateClearCon.Execute)

	// /cocktails
	r.GET("/cocktails/individual/:id", findCocktailCon.Execute)          // 認証なしでカクテル取得
	authRouter.GET("/cocktails/individual/:id", findCocktailCon.Execute) // 認証ありでカクテル取得
	r.GET("/cocktails/list", findCocktailListCon.Execute)                // 認証なしでカクテル一覧取得
	authRouter.GET("/cocktails/list", findCocktailListCon.Execute)       // 認証ありでカクテル一覧取得

	// /cocktail_images
	authRouter.POST("/cocktail_images", saveCocktailImageCon.Execute)

	r.Run()
}
