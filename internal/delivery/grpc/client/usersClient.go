package client

import (
	"context"
	proto "github.com/DoktorGhost/external-api/src/go/pkg/grpc/clients/api/grpc/protobuf/clients_v1"
	"github.com/DoktorGhost/golibrary/config"
	"github.com/DoktorGhost/golibrary/internal/core/user/entities"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

type UserClient struct {
	proto.ClientsServiceClient
}

func InitUserClient() (*UserClient, *grpc.ClientConn) {
	// Подключаемся к gRPC-серверу USER
	conn, err := grpc.Dial(config.LoadConfig().GrpcConfig.UserHost+":"+config.LoadConfig().GrpcConfig.UserPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// Создаем gRPC-клиента для сервиса USER
	userClient := proto.NewClientsServiceClient(conn)

	// Создаем сервис User, который будет использовать этот клиент
	userService := &UserClient{userClient}

	log.Println("Connected to User service port:", config.LoadConfig().GrpcConfig.UserPort)
	return userService, conn
}

func (a *UserClient) Register(userData entities.RegisterData) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	regData := &proto.RegisterRequest{
		Username:   userData.Username,
		Password:   userData.Password,
		Name:       userData.Name,
		Surname:    userData.Surname,
		Patronymic: userData.Patronymic,
	}
	// Вызываем метод Register в сервисе USER
	resp, err := a.ClientsServiceClient.Register(ctx, regData)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(resp.Id, 10), nil
}

func (a *UserClient) Login(userData entities.Login) (entities.UserTable, error) {
	var result entities.UserTable
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Вызываем метод Login в сервисе USER
	resp, err := a.ClientsServiceClient.Login(ctx, &proto.LoginRequest{Username: userData.Username, Password: userData.Password})
	if err != nil {
		return result, err
	}

	result.ID = int(resp.Id)
	result.Username = resp.Username
	result.PasswordHash = resp.Password
	result.FullName = resp.Fullname

	return result, nil
}

func (a *UserClient) GetUserById(userID int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Вызываем метод GetUserById в сервисе USER
	resp, err := a.ClientsServiceClient.GetUserByID(ctx, &proto.UserID{Id: int64(userID)})
	if err != nil {
		return "", err
	}

	return resp.Username, nil
}
