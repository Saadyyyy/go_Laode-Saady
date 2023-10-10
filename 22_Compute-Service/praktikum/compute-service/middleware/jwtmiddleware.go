package middleware

import (
	"praktikum_22/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func PasswordVerify(storedPassoword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassoword), []byte(password))
	return err == nil
}
