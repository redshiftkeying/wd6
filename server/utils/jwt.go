package utils

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// make Token
func GetToken(username string, role string) (*string, error) {
	claims := make(jwt.MapClaims)
	claims["username"] = username
	claims["role"] = role
	// exp default 18 hour config from system setting
	exp := 18
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(exp)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // TODO: viper Signing Method
	sal := "TJVA95OrM4E2cBab30RMHrRD1EfxjoYZge2ONFh7Hg9"       // TODO: viper
	tokenString, err := token.SignedString([]byte(sal))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &tokenString, nil
}

// CheckToken
func CheckToken(key string) bool {
	kv := strings.Split(key, " ")
	// Key type check. the Bearer token.
	if len(kv) != 2 || kv[0] != "Bearer" {
		fmt.Println("Token invalid:", key)
		return false
	}
	tokenString := kv[1]
	sal := "weqweqweweweqwdeqwqdwdw"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(sal), nil
	})
	if err != nil {
		log.Println("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return false
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return false
			} else {
				// Couldn't handle this token
				return false
			}
		} else {
			// Couldn't handle this token
			return false
		}
	}
	if !token.Valid {
		log.Println("Token invalid:", tokenString)
		return false
	}
	return true
}
