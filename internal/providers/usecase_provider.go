package providers

import (
	domainBook "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/usecases"
	domainUser "github.com/DoktorGhost/golibrary/internal/core/user/usecases"
)

type UseCaseProvider struct {
	BookUseCase    *domainBook.BookUseCase
	LibraryUseCase *domainBook.LibraryUseCase
	DataUseCase    *domainUser.DataUseCase
	UserUseCase    *domainUser.UsersUseCase
	AuthUseCase    *domainUser.AuthUseCase
}

func NewUseCaseProvider() *UseCaseProvider {
	return &UseCaseProvider{}
}

func (ucp *UseCaseProvider) RegisterDependencies(provider *ServiceProvider, key string) {
	ucp.BookUseCase = domainBook.NewBookUseCase(provider.bookService, provider.authorService, provider.rentalService)
	ucp.LibraryUseCase = domainBook.NewLibraryUseCase(provider.rentalService, provider.usersService, ucp.BookUseCase, provider.authorService)
	ucp.UserUseCase = domainUser.NewUsersUseCase(provider.usersService)
	ucp.DataUseCase = domainUser.NewDataUseCase(provider.bookService, provider.rentalService, provider.authorService, ucp.UserUseCase)
	ucp.AuthUseCase = domainUser.NewAuthUseCase(provider.usersService, key)
}
