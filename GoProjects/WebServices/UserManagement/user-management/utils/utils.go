package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"regexp"
	"time"
)

// ValidateEmail 验证邮箱格式
func ValidateEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}

// ValidatePassword 验证密码强度
func ValidatePassword(password string) bool {
	return len(password) >= 8 // 至少8位
}

// GenerateJWT 生成JWT Token
func GenerateJWT(userID, roleID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role_id": roleID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte("supersecretkey"))
}

// ValidateJWT 验证JWT Token
func ValidateJWT(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("supersecretkey"), nil
	})
	return err == nil && token.Valid
}
