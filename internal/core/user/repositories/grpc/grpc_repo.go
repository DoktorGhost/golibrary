package grpcrepo

import (
	"github.com/DoktorGhost/golibrary/internal/core/user/entities"
	"github.com/DoktorGhost/golibrary/internal/delivery/grpc/client"
)

type UsersRepository struct {
	client *client.UserClient
}

func NewUsersRepository(client *client.UserClient) *UsersRepository {
	return &UsersRepository{client: client}
}

func (ur *UsersRepository) Register(data entities.RegisterData) (string, error) {
	return ur.client.Register(data)
}

func (ur *UsersRepository) Login(data entities.Login) (entities.UserTable, error) {
	return ur.client.Login(data)
}

func (ur *UsersRepository) GetUserById(userID int) (string, error) {
	return ur.client.GetUserById(userID)
}
