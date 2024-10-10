package app

import (
	"database/sql"
	"github.com/DoktorGhost/golibrary/config"
	"sync"

	"github.com/DoktorGhost/golibrary/internal/providers"
)

var (
	Container container
	once      sync.Once
)

type container struct {
	UseCaseProvider *providers.UseCaseProvider
}

func Init(db *sql.DB, conf *config.Config) container {
	once.Do(func() {
		repositoryProvider := providers.NewRepositoryProvider(db)
		repositoryProvider.RegisterDependencies()

		serviceProvider := providers.NewServiceProvider()
		serviceProvider.RegisterDependencies(repositoryProvider)

		useCaseProvider := providers.NewUseCaseProvider()
		useCaseProvider.RegisterDependencies(serviceProvider, conf.JWTSecret)

		Container = container{
			UseCaseProvider: useCaseProvider,
		}
	})
	return Container
}
