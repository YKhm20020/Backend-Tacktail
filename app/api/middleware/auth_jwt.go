package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// cookieからjwtを取得
		tokenString, err := ctx.Cookie("jwt")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   errors.New("unauthorized").Error(),
				"message": "cookieからjwtの検出に失敗",
			})
			ctx.Abort()
			return
		}

		// jwtの検証
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			secretKey := os.Getenv("SECRET_KEY")
			return []byte(secretKey), nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   errors.New("unauthorized").Error(),
				"message": "jwtの検証に失敗",
			})
			ctx.Abort()
			return
		}

		// トークンからユーザー名を取得し、contextにセット
		userName, err := extractUserName(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   errors.New("invalid token").Error(),
				"message": "トークンからユーザー名の取得に失敗",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user", userName)

		// 認証を突破!!
		ctx.Next()
	}
}

func extractUserName(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", nil
	}

	userName, ok := claims["user_name"].(string)

	if !ok {
		return "", nil
	}
	return userName, nil
}
