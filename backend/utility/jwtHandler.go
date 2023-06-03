package utility

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// Token Generation
func GenerateJWTToken(userType string, license int, exp time.Duration) (tokenString string, err error) {
	// Generating the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userType": userType,
		"license":  license,
		"exp":      time.Now().Add(exp).Unix(),
	})

	// Signing it
	tokenString, err = token.SignedString([]byte(os.Getenv("SECRET")))
	return
}
