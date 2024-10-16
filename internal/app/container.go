package app

import (
	"github.com/DoktorGhost/golibrary/config"
	"github.com/DoktorGhost/golibrary/internal/providers"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Container container
)

type container struct {
	UseCaseProvider *providers.UseCaseProvider
}

func Init(db *pgxpool.Pool) container {
	repositoryProvider := providers.NewRepositoryProvider(db)
	repositoryProvider.RegisterDependencies()

	serviceProvider := providers.NewServiceProvider()
	serviceProvider.RegisterDependencies(repositoryProvider)

	useCaseProvider := providers.NewUseCaseProvider()
	useCaseProvider.RegisterDependencies(serviceProvider, config.LoadConfig().Secrets.JWTSecret)

	Container = container{
		UseCaseProvider: useCaseProvider,
	}

	return Container
}
