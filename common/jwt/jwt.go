package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/t1tc01/test-directive/model"
)

var (
	SecretKey = []byte("secret")
)

type ParseResult struct {
	Username   string
	Role       model.Role
	ExpireTime time.Time
}

func GenerateToken(username string, role model.Role) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil

}

func ParseToken(tokenStr string) (ParseResult, error) {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		role := claims["role"].(string)
		exp := claims["exp"].(float64)

		var roleUser model.Role = model.Role(role)
		var expTime = time.Unix(int64(exp), 0)

		result := ParseResult{Username: username, Role: roleUser, ExpireTime: expTime}
		return result, nil
	} else {
		return ParseResult{}, err
	}
}
