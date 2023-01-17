package memory

import (
	"context"

	"github.com/google/uuid"

	models "github.com/Nikolay-Yakushev/mango/internal/domain"

	"github.com/Nikolay-Yakushev/mango/internal/domain/entities/users"
)

func (imem *InMemoryStorage) GetUser(
	ctx context.Context, login string) (users.User, error) {
	
	user, ok := imem.storage[login]
	if !ok {
		imem.log.Sugar().Errorw("user not found", "login", login)
		err := models.NotFoundErr
		return users.User{}, err
	}
	return user, nil
}

func (imem *InMemoryStorage) BlockUser (ctx context.Context, u users.User) (error){
	user, ok := imem.block[u.Login]
	if ok{
		imem.log.Sugar().Errorw("user already blocked", "login", u.Login)
		err := models.ConflictErr
		return err
	}
	imem.block[u.Login] = user
	return nil
} 

func (imem *InMemoryStorage) SetUser(
	ctx context.Context, login, password, email string) (users.User, error) {

	user, ok := imem.storage[login]
	
	if ok {
		imem.log.Sugar().Errorw("user already exists", "login", login)
		err := models.ConflictErr
		return user, err
	}

	genId := func()uuid.UUID {
		id := uuid.New() 
		return id
	}

	u := users.User{
		Id:       genId(),
		Login:    login,
		Password: password,
		Email:    email,
	}

	hash, err := u.HashPassword()
	if err != nil{
		imem.log.Sugar().Errorw("Failed to hash password", "reason", err)
		return users.User{}, err
	}
	u.Password = hash

	imem.storage[login] = u
	return u, nil


}
