package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ヘッダーからjwtを取得
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			// ヘッダーのAuthorizationが空
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   errors.New("unauthorized").Error(),
				"message": "ヘッダー内のAuthorizationがない",
			})
			ctx.Abort()
			return
		}

		contents := strings.Split(authorizationHeader, " ")
		if len(contents) != 2 || contents[0] == "Bearer" {
			// AuthorizationにBearerが存在しない
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   errors.New("unauthorized").Error(),
				"message": "AuthorizationにBearerが存在しない",
			})
			ctx.Abort()
			return
		}

		// jwtの検証
		token, err := jwt.Parse(contents[1], func(t *jwt.Token) (interface{}, error) {
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
		userName, userID, err := extractUser(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   errors.New("invalid token").Error(),
				"message": "トークンからユーザー名およびユーザーIDの取得に失敗",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user_name", userName)
		ctx.Set("user_id", userID)

		// 認証を突破!!
		ctx.Next()
	}
}

func extractUser(token *jwt.Token) (string, string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", "", nil
	}

	userName, ok := claims["user_name"].(string)
	if !ok {
		return "", "", nil
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", "", nil
	}

	return userName, userID, nil
}
