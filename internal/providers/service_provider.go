package providers

import (
	domainBook "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/services"
	domainRental "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/services"
	domainUser "github.com/DoktorGhost/golibrary/internal/core/user/services"
)

type ServiceProvider struct {
	authorService *domainBook.AuthorService
	bookService   *domainBook.BookService
	rentalService *domainRental.RentalService
	usersService  *domainUser.UserService
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) RegisterDependencies(provider *RepositoryProvider) {
	s.authorService = domainBook.NewAuthorService(provider.bookRepositoryPostgres)
	s.bookService = domainBook.NewBookService(provider.bookRepositoryPostgres)
	s.rentalService = domainRental.NewRentalService(provider.rentalRepositoryPostgres)
	s.usersService = domainUser.NewUserService(provider.usersRepositoryPostgres)
}
