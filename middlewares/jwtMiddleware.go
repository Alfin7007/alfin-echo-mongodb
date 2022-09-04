package middlewares

import (
	"errors"
	"explore/mongodb/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMIddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.JWTKey()),
	})
}

func CreateToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(12 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTKey()))
}

func ExtracToken(e echo.Context) (string, error) {
	token := e.Get("user")
	if token == nil {
		return "", errors.New("not authorized")
	}
	user := token.(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userID := claims["userID"].(string)
		return userID, nil
	}
	return "", errors.New("invalid token")
}
