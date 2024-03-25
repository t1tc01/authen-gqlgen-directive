package users

import (
	"context"

	"github.com/t1tc01/test-directive/common"
	"golang.org/x/crypto/bcrypt"
)

// CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Authenticate Login
func AuthenticateLogin(c context.Context, username string, password string) bool {

	hash := ""
	for _, user := range common.UserSample {
		if user.Username == username {
			hash = user.HashPassword
		}
	}

	return CheckPasswordHash(password, hash)
}
