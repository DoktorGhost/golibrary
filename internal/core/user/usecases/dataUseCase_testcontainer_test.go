package usecases

import (
	bookRepo "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres"
	service2 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/services"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/usecases"
	rentalRepo "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/repositories/postgres"
	service3 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/services"
	userRepo "github.com/DoktorGhost/golibrary/internal/core/user/repositories/postgres"
	"github.com/DoktorGhost/golibrary/internal/core/user/services"
	"github.com/DoktorGhost/golibrary/pkg/storage/test_container"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddLibrary2(t *testing.T) {
	db, cl := test_container.SetupPostgresContainer(t)
	defer cl()

	//моки
	repoUser := userRepo.NewPostgresRepository(db)
	repoBook := bookRepo.NewBookPostgresRepository(db)
	repoRental := rentalRepo.NewRentalPostgresRepository(db)

	//сервисы
	userService := services.NewUserService(repoUser)
	bookService := service2.NewBookService(repoBook)
	rentalService := service3.NewRentalService(repoRental)
	authorService := service2.NewAuthorService(repoBook)

	//юзкейсы
	userUseCase := NewUsersUseCase(userService)
	bookUseCase := usecases.NewBookUseCase(bookService, authorService, rentalService)
	libraryUseCase := usecases.NewLibraryUseCase(rentalService, userService, bookUseCase, authorService)
	dataUseCase := NewDataUseCase(bookService, rentalService, authorService, userUseCase)

	err := dataUseCase.AddLibrary()

	if err != nil {
		t.Fatal(err)
	}

	var count int

	// 1. Проверка количества добавленных авторов
	t.Run("Проверка количества добавленных авторов", func(t *testing.T) {
		err = db.QueryRow(`SELECT COUNT(*) FROM library.authors;`).Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 10, count)
	})

	// 2. Проверка количества добавленных книг
	t.Run("Проверка количества добавленных книг", func(t *testing.T) {
		err = db.QueryRow(`SELECT COUNT(*) FROM library.books;`).Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 100, count)
	})

	// 3. Проверка количества записей в library.rentals
	t.Run("Проверка количества записей в library.rentals", func(t *testing.T) {
		err = db.QueryRow(`SELECT COUNT(*) FROM library.rentals;`).Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 100, count)
	})

	// 4. Проверка количества добавленных пользователей
	t.Run("Проверка количества добавленных пользователей", func(t *testing.T) {
		err = db.QueryRow(`SELECT COUNT(*) FROM users.users;`).Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 60, count)
	})

	// 5. Проверка создания записи аренды книги
	t.Run("Проверка аренды книги", func(t *testing.T) {
		id, err := libraryUseCase.GiveBook(1, 1)

		assert.NoError(t, err)
		assert.Equal(t, 1, id)

		err = db.QueryRow(`SELECT COUNT(*) FROM library.rentals_info;`).Scan(&count)

		assert.NoError(t, err)
		assert.Equal(t, 1, count)

		var rentalsID int
		err = db.QueryRow(`SELECT rentals_id FROM library.rentals WHERE id=1;`).Scan(&rentalsID)
		assert.NoError(t, err)
		assert.Equal(t, 1, rentalsID)

		err = libraryUseCase.BackBook(1)
		assert.NoError(t, err)

		err = db.QueryRow(`SELECT rentals_id FROM library.rentals WHERE id=1;`).Scan(&rentalsID)
		assert.Error(t, err)
	})

}
