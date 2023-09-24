package middleware

import (
	"Laode_Saady/18_Middleware/praktikum/constant"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func CreateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_JWT))
}

func JwtMiddleware() echo.MiddlewareFunc {
	return echoJwt.WithConfig(echoJwt.Config{
		SigningKey:    []byte(constant.SECRET_JWT),
		SigningMethod: "HS256",
	})
}
