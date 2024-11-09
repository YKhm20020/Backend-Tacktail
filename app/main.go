package main

import (
	"github.com/YKhm20020/Backend-Tacktail/api"
	"github.com/YKhm20020/Backend-Tacktail/infrastructure"
	_ "github.com/lib/pq"
)

func main() {
	// データベース接続
	db := infrastructure.SetupDB()

	// ルーティング設定
	api.SetupRouter(db)
}
