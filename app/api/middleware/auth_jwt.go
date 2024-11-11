package middleware

import "github.com/golang-jwt/jwt/v5"

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
