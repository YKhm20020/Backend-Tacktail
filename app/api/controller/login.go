package controller

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/YKhm20020/Backend-Tacktail/usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Login struct {
	uc usecase.Login
}

func NewLogin(uc usecase.Login) Login {
	return Login{
		uc: uc,
	}
}

func (con Login) Execute(ctx *gin.Context) {
	var input usecase.LoginInput

	// jsonデータを構造体にデコードしバインド
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ログイン処理に成功したか判定
	user, err := con.uc.Execute(input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// JWTの発行
	claims := jwt.MapClaims{
		"user_id":   user.Id,
		"user_name": user.Name,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": errors.New("failed to generate token").Error()},
		)
		return
	}

	// cookie認証じゃなくなったのでいらない
	// ctx.SetSameSite((http.SameSiteNoneMode))

	// cookieにjwtをセット
	// if product := os.Getenv(("PRODUCTION")); product == "" {
	// 	// 開発用サーバー
	// 	ctx.SetCookie("jwt", tokenString, 3600, "/", "localhost", false, true)
	// } else {
	// 	// APIサーバーのデプロイ先
	// 	ctx.SetCookie("jwt", tokenString, 3600, "/", "", true, true)
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"id":    user.Id,
		"name":  user.Name,
		"token": tokenString,
	})
}
