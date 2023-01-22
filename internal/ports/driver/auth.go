package ports

import (
	"context"

	"github.com/Nikolay-Yakushev/mango/internal/domain/entities/users"
)

type Auth interface {
	Login(ctx context.Context, login, password string)(string, string, error)
	Singup(ctx context.Context, login, password, email string)(users.User, error)
	Logout(ctx context.Context, login, password string)(bool, error)
	Verify(ctx context.Context, login, password string)(users.VerifyResponse, error)
}