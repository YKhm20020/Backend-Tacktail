package main

import (
	"github.com/YKhm20020/Backend-Tacktail/api"
	"github.com/YKhm20020/Backend-Tacktail/infrastructure"
	_ "github.com/lib/pq"
)

// @title gin-swagger todos
// @version 1.0
// @license.name kosuke
// @description このswaggerはgin-swaggerの見本apiです
func main() {
	// データベース接続
	db := infrastructure.SetupDB()

	api.SetupRouter(db)
}
