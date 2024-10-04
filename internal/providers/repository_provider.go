package providers

import (
	"database/sql"

	subdomainBook "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres"
	subdomainRental "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/repositories/postgres"
	domainUser "github.com/DoktorGhost/golibrary/internal/core/user/repositories/postgres"
)

type RepositoryProvider struct {
	db *sql.DB

	bookRepositoryPostgres   *subdomainBook.BookRepository
	rentalRepositoryPostgres *subdomainRental.RentalRepository
	usersRepositoryPostgres  *domainUser.UsersRepository
}

func NewRepositoryProvider(db *sql.DB) *RepositoryProvider {
	return &RepositoryProvider{db: db}
}

func (r *RepositoryProvider) RegisterDependencies() {
	r.bookRepositoryPostgres = subdomainBook.NewBookPostgresRepository(r.db)
	r.rentalRepositoryPostgres = subdomainRental.NewRentalPostgresRepository(r.db)
	r.usersRepositoryPostgres = domainUser.NewPostgresRepository(r.db)
}
