package providers

import (
	grpcrepo "github.com/DoktorGhost/golibrary/internal/core/user/repositories/grpc"
	"github.com/DoktorGhost/golibrary/internal/delivery/grpc/client"
	"github.com/jackc/pgx/v5/pgxpool"

	subdomainBook "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres"
	subdomainRental "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/repositories/postgres"
)

type RepositoryProvider struct {
	db                       *pgxpool.Pool
	client                   *client.UserClient
	bookRepositoryPostgres   *subdomainBook.BookRepository
	rentalRepositoryPostgres *subdomainRental.RentalRepository
	usersRepositoryPostgres  *grpcrepo.UsersRepository
}

func NewRepositoryProvider(db *pgxpool.Pool, client *client.UserClient) *RepositoryProvider {
	return &RepositoryProvider{db: db, client: client}
}

func (r *RepositoryProvider) RegisterDependencies() {
	r.bookRepositoryPostgres = subdomainBook.NewBookPostgresRepository(r.db)
	r.rentalRepositoryPostgres = subdomainRental.NewRentalPostgresRepository(r.db)
	r.usersRepositoryPostgres = grpcrepo.NewUsersRepository(r.client)
}
