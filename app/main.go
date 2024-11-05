package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/KakinokiKanta/go-dev-template/api/controller"
	"github.com/KakinokiKanta/go-dev-template/infrastructure/repository"
	"github.com/KakinokiKanta/go-dev-template/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// データベース接続
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"db", "db-user", "db-password", "db-name", "5432")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userRepository := repository.NewUserRepository(db)
	createUserUc := usecase.NewCreateUser(userRepository)
	createUserCon := controller.NewCreateUser(createUserUc)

	r.POST("/users", createUserCon.Execute)
	r.Run()
}
