package db

import (
	"github.com/t1tc01/test-directive/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string
	HashPassword string
	Role         model.Role
}

func CreateUserSample() ([]User, error) {
	userArray := make([]User, 1)

	adminPassword := "admin"
	user1Password := "user1"

	adminHashPassword, _ := HashPassword(adminPassword)
	user1HashPassword, _ := HashPassword(user1Password)

	userArray = append(userArray, User{Username: "admin", HashPassword: adminHashPassword, Role: model.RoleAdmin})
	userArray = append(userArray, User{Username: "user1", HashPassword: user1HashPassword, Role: model.RoleUser})

	return userArray, nil
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
