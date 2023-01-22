package users

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	Login    string
	Password string
	Email    string
}

type VerifyResponse struct {
	AccessToken  string
	RefreshToken string
	User         User
}