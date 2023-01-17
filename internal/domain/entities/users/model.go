package users

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uuid.UUID
	Login    string
	Password string
	Email    string
}

func (u User) HashPassword()(string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
		return "", err
    }
    return string(hash), nil
}

type VerifyResponse struct {
	AccessToken  string
	RefreshToken string
	User         User
}