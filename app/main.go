package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	id int
	name string
	password string
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

	r.GET("/get", func(ctx *gin.Context) {
		row := db.QueryRow("SELECT id, name, password FROM accounts;")
		if row.Err() != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": row.Err().Error()})
			return
		}
		err := row.Scan(&id, &name, &password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"id": strconv.Itoa(id),
			"name": name,
			"password": password,
		})
	})
	r.Run()
}
