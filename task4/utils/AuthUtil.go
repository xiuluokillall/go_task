package utils

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

/*
加密密码
*/
func GenerateFromPassword(password string) (string, error) {
	passBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passBytes), err
}

// 验证密码
func CompareHashAndPassword(origin string, target string) error {
	return bcrypt.CompareHashAndPassword([]byte(origin), []byte(target))
}

/*
*
生成JWT
*/
func NewWithClaims(userId uint, username string) (string, error) {
	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userId,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte("your_secret_key"))
}
