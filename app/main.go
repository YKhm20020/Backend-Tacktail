package main

import (
	"github.com/KakinokiKanta/go-dev-template/api"
	"github.com/KakinokiKanta/go-dev-template/infrastructure"
	_ "github.com/lib/pq"
)

func main() {
	// データベース接続
	db := infrastructure.SetupDB()

	api.SetupRouter(db)
}
