package main

import (
	"github.com/YKhm20020/Backend-Tacktail/api"
	"github.com/YKhm20020/Backend-Tacktail/docs"
	"github.com/YKhm20020/Backend-Tacktail/infrastructure"
	_ "github.com/lib/pq"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title gin-swagger todos
// @version 1.0
// @license.name kosuke
// @description このswaggerはgin-swaggerの見本apiです
func main() {
	// データベース接続
	db := infrastructure.SetupDB()

	router := api.SetupRouter(db)

	// swagger用エンドポイント
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()
}
