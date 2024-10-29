package providers

import (
	grpcBook "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/grpc"
	grpcUser "github.com/DoktorGhost/golibrary/internal/core/user/repositories/grpc"
	"github.com/DoktorGhost/golibrary/internal/delivery/grpc/client"
	"github.com/jackc/pgx/v5/pgxpool"

	subdomainRental "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/repositories/postgres"
)

type RepositoryProvider struct {
	db                       *pgxpool.Pool
	userClient               *client.UserClient
	bookClient               *client.BookClient
	rentalRepositoryPostgres *subdomainRental.RentalRepository
	usersRepositoryPostgres  *grpcUser.UsersRepository
	bookRepositoryPostgres   *grpcBook.BookRepository
}

func NewRepositoryProvider(db *pgxpool.Pool, userClient *client.UserClient, bookClient *client.BookClient) *RepositoryProvider {
	return &RepositoryProvider{db: db, userClient: userClient, bookClient: bookClient}
}

func (r *RepositoryProvider) RegisterDependencies() {
	r.rentalRepositoryPostgres = subdomainRental.NewRentalPostgresRepository(r.db)
	r.usersRepositoryPostgres = grpcUser.NewUsersRepository(r.userClient)
	r.bookRepositoryPostgres = grpcBook.NewBookRepository(r.bookClient)
}
