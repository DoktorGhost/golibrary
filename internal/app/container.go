package app

import (
	"github.com/DoktorGhost/golibrary/config"
	"github.com/DoktorGhost/golibrary/internal/delivery/grpc/client"
	"github.com/DoktorGhost/golibrary/internal/providers"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

var (
	Container container
	once      sync.Once
)

type container struct {
	UseCaseProvider *providers.UseCaseProvider
}

func Init(db *pgxpool.Pool, userClient *client.UserClient, bookClient *client.BookClient) container {
	once.Do(func() {
		repositoryProvider := providers.NewRepositoryProvider(db, userClient, bookClient)
		repositoryProvider.RegisterDependencies()

		serviceProvider := providers.NewServiceProvider()
		serviceProvider.RegisterDependencies(repositoryProvider)

		useCaseProvider := providers.NewUseCaseProvider()
		useCaseProvider.RegisterDependencies(serviceProvider, config.LoadConfig().Secrets.JWTSecret)

		Container = container{
			UseCaseProvider: useCaseProvider,
		}
	})

	return Container
}
