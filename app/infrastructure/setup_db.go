package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	host     string
	user     string
	password string
	dbname   string
	port     string
)

func SetupDB() *sql.DB {
	err := loadEnv()
	if err != nil {
		panic("完了変数の取得に失敗")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}

	return db
}

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	host = os.Getenv("DB_HOST")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
	port = os.Getenv("DB_PORT")

	return nil
}
