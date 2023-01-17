package memory

import (
	"go.uber.org/zap"
	"github.com/Nikolay-Yakushev/mango/internal/domain/entities/users"
)

type InMemoryStorage struct {
	storage map[string]users.User
	block   map[string]users.User
	log     *zap.Logger
}

func New(logger *zap.Logger) (*InMemoryStorage, error) {
	storage := make(map[string]users.User)
	block := make(map[string]users.User)

	return &InMemoryStorage{
		storage: storage,
		log: logger,
		block: block}, nil
}
