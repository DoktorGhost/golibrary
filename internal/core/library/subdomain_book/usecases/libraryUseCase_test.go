package usecases

import (
	"errors"
	dao3 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
	service2 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/services"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/entities"
	dao2 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/repositories/postgres/dao"

	service3 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/services"
	"github.com/DoktorGhost/golibrary/internal/core/user/repositories/postgres/dao"
	"github.com/DoktorGhost/golibrary/internal/core/user/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLibraryUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//моки
	mockRepoRental := service3.NewMockRentalsRepository(ctrl)
	mockRepoUser := services.NewMockUserRepository(ctrl)
	mockRepoBook := service2.NewMockBookRepository(ctrl)
	mockRepoAuthor := service2.NewMockAuthorRepository(ctrl)

	//сервисы
	rentalService := service3.NewRentalService(mockRepoRental)
	userService := services.NewUserService(mockRepoUser)
	authorService := service2.NewAuthorService(mockRepoAuthor)
	bookService := service2.NewBookService(mockRepoBook)

	//юзкейсы
	//userUseCase := usecases.NewUsersUseCase(userService)
	bookUseCase := NewBookUseCase(bookService, authorService, rentalService)
	libraryUseCase := NewLibraryUseCase(rentalService, userService, bookUseCase, authorService)

	// 1. GiveBook. Ошибка в GetUserById
	t.Run("Ошибка в GetUserById", func(t *testing.T) {
		bookID := 1
		userID := 1
		mockRepoUser.EXPECT().GetUserByID(userID).Return(dao.UserTable{}, errors.New("ERROR"))

		rentalID, err := libraryUseCase.GiveBook(bookID, userID)

		assert.Error(t, err)
		assert.Equal(t, 0, rentalID)
	})

	// 2. GiveBook. Ошибка в GetRentalsByID
	t.Run("Ошибка в GetRentalsByID", func(t *testing.T) {
		bookID := 1
		userID := 1
		mockRepoUser.EXPECT().GetUserByID(userID).Return(dao.UserTable{}, nil)
		mockRepoRental.EXPECT().GetRentalsByID(bookID).Return(0, errors.New("ERROR"))

		rentalID, err := libraryUseCase.GiveBook(bookID, userID)

		assert.Error(t, err)
		assert.Equal(t, 0, rentalID)
	})

	// 3. GiveBook. rentalID > 0
	t.Run("rentalID > 0", func(t *testing.T) {
		bookID := 1
		userID := 1
		mockRepoUser.EXPECT().GetUserByID(userID).Return(dao.UserTable{}, nil)
		mockRepoRental.EXPECT().GetRentalsByID(bookID).Return(2, nil)

		rentalID, err := libraryUseCase.GiveBook(bookID, userID)

		assert.Error(t, err)
		assert.Equal(t, 0, rentalID)
	})

	// 4. GiveBook. CreateRentalInfo fail
	t.Run(" CreateRentalInfo fail", func(t *testing.T) {
		bookID := 1
		userID := 1
		mockRepoUser.EXPECT().GetUserByID(userID).Return(dao.UserTable{}, nil)
		mockRepoRental.EXPECT().GetRentalsByID(bookID).Return(0, nil)
		mockRepoRental.EXPECT().CreateRentalsInfo(userID, bookID).Return(0, errors.New("ERROR"))

		rentalID, err := libraryUseCase.GiveBook(bookID, userID)

		assert.Error(t, err)
		assert.Equal(t, 0, rentalID)
	})

	// 5. GiveBook. UpdateRentals fail
	t.Run("UpdateRentals fail", func(t *testing.T) {
		bookID := 1
		userID := 1
		rentalID := 1
		mockRepoUser.EXPECT().GetUserByID(userID).Return(dao.UserTable{}, nil)
		mockRepoRental.EXPECT().GetRentalsByID(bookID).Return(0, nil)
		mockRepoRental.EXPECT().CreateRentalsInfo(userID, bookID).Return(rentalID, nil)
		mockRepoRental.EXPECT().UpdateRentals(bookID, rentalID).Return(errors.New("ERROR"))

		rentalID, err := libraryUseCase.GiveBook(bookID, userID)

		assert.Error(t, err)
		assert.Equal(t, 0, rentalID)
	})

	// 5. GiveBook success
	t.Run("GiveBook success", func(t *testing.T) {
		bookID := 1
		userID := 1
		rentalID := 1
		mockRepoUser.EXPECT().GetUserByID(userID).Return(dao.UserTable{}, nil)
		mockRepoRental.EXPECT().GetRentalsByID(bookID).Return(0, nil)
		mockRepoRental.EXPECT().CreateRentalsInfo(userID, bookID).Return(rentalID, nil)
		mockRepoRental.EXPECT().UpdateRentals(bookID, rentalID).Return(nil)

		rentalID, err := libraryUseCase.GiveBook(bookID, userID)

		assert.NoError(t, err)
		assert.Equal(t, 1, rentalID)
	})

	// 5. BackBook. GetRentalsByID fail
	t.Run("GetRentalsByID fail", func(t *testing.T) {
		bookID := 1

		mockRepoRental.EXPECT().GetRentalsByID(bookID).Return(1, errors.New("Error"))

		err := libraryUseCase.BackBook(bookID)

		assert.Error(t, err)
	})

	// 6. BackBook. rentalID == 0
	t.Run("rentalID == 0", func(t *testing.T) {
		bookID := 1

		mockRepoRental.EXPECT().GetRentalsByID(bookID).Return(0, nil)

		err := libraryUseCase.BackBook(bookID)

		assert.Error(t, err)
	})

	// 7. BackBook. UpdateRentals fail
	t.Run("UpdateRentals fail", func(t *testing.T) {
		bookID := 1
		mockRepoRental.EXPECT().GetRentalsByID(bookID).Return(2, nil)
		mockRepoRental.EXPECT().UpdateRentals(bookID, 0).Return(errors.New("Error"))

		err := libraryUseCase.BackBook(bookID)

		assert.Error(t, err)
	})

	// 8. BackBook. GetRentalsInfoByID fail
	t.Run("GetRentalsInfoByID fail", func(t *testing.T) {
		bookID := 1
		rentalID := 2
		mockRepoRental.EXPECT().GetRentalsByID(bookID).Return(rentalID, nil)
		mockRepoRental.EXPECT().UpdateRentals(bookID, 0).Return(nil)
		mockRepoRental.EXPECT().GetRentalsInfoByID(rentalID).Return(dao2.RentalsTable{}, errors.New("Error"))

		err := libraryUseCase.BackBook(bookID)

		assert.Error(t, err)
	})

	// 9. BackBook. UpdateRentalsInfo fail
	t.Run("UpdateRentalsInfo fail", func(t *testing.T) {
		bookID := 1
		rentalID := 2
		rental := dao2.RentalsTable{}
		mockRepoRental.EXPECT().GetRentalsByID(bookID).Return(rentalID, nil)
		mockRepoRental.EXPECT().UpdateRentals(bookID, 0).Return(nil)
		mockRepoRental.EXPECT().GetRentalsInfoByID(rentalID).Return(rental, nil)
		mockRepoRental.EXPECT().UpdateRentalsInfo(gomock.Any()).Return(errors.New("Error"))

		err := libraryUseCase.BackBook(bookID)

		assert.Error(t, err)
	})

	// 10. BackBook Success
	t.Run("BackBook Success", func(t *testing.T) {
		bookID := 1
		rentalID := 2
		rental := dao2.RentalsTable{}
		mockRepoRental.EXPECT().GetRentalsByID(bookID).Return(rentalID, nil)
		mockRepoRental.EXPECT().UpdateRentals(bookID, 0).Return(nil)
		mockRepoRental.EXPECT().GetRentalsInfoByID(rentalID).Return(rental, nil)
		mockRepoRental.EXPECT().UpdateRentalsInfo(gomock.Any()).Return(nil)

		err := libraryUseCase.BackBook(bookID)

		assert.NoError(t, err)
	})

	//11. GetUserRentals. GetActiveRentals fail
	t.Run("GetActiveRentals fail", func(t *testing.T) {
		var result []entities.UserWithRentedBooks
		mockRepoRental.EXPECT().GetActiveRentals().Return(nil, errors.New("Error"))

		userWithRentedBooks, err := libraryUseCase.GetUserRentals()

		assert.Error(t, err)
		assert.Equal(t, result, userWithRentedBooks)
	})

	//12. GetUserRentals. len(rentalsID) < 1
	t.Run("len(rentalsID) < 1", func(t *testing.T) {
		var result map[int][]int
		mockRepoRental.EXPECT().GetActiveRentals().Return(result, nil)

		_, err := libraryUseCase.GetUserRentals()

		assert.Error(t, err)
	})

	//13. GetUserRentals. GetUserById fail
	t.Run("GetUserById fail", func(t *testing.T) {
		result := make(map[int][]int)
		result[1] = []int{1, 2, 3, 4}
		result[2] = []int{1, 2, 3, 4}

		mockRepoRental.EXPECT().GetActiveRentals().Return(result, nil)
		mockRepoUser.EXPECT().GetUserByID(1).Return(dao.UserTable{}, errors.New("Error"))

		_, err := libraryUseCase.GetUserRentals()

		assert.Error(t, err)
	})

	//14. GetUserRentals success
	t.Run("GetUserRentals success", func(t *testing.T) {
		result := make(map[int][]int)
		result[1] = []int{1, 2}
		result[2] = []int{3, 4}

		mockRepoRental.EXPECT().GetActiveRentals().Return(result, nil)

		// Ожидаем, что GetUserByID будет вызван дважды (для пользователей 1 и 2)
		mockRepoUser.EXPECT().GetUserByID(1).Return(dao.UserTable{}, nil)
		mockRepoUser.EXPECT().GetUserByID(2).Return(dao.UserTable{}, nil)

		// Ожидаем, что GetBookByID будет вызван для книг 1, 2, 3 и 4
		mockRepoBook.EXPECT().GetBookByID(1).Return(dao3.BookTable{}, nil)
		mockRepoBook.EXPECT().GetBookByID(2).Return(dao3.BookTable{}, nil)
		mockRepoBook.EXPECT().GetBookByID(3).Return(dao3.BookTable{}, nil)
		mockRepoBook.EXPECT().GetBookByID(4).Return(dao3.BookTable{}, nil)

		// Ожидаем, что GetAuthorByID будет вызван для авторов книг 1, 2, 3 и 4
		mockRepoAuthor.EXPECT().GetAuthorByID(gomock.Any()).Return(dao3.AuthorTable{}, nil).Times(4)

		_, err := libraryUseCase.GetUserRentals()

		assert.NoError(t, err)
	})

	//15. GetTopAuthors GetTopAuthorsByPeriod fail
	t.Run("GetTopAuthorsByPeriod fail", func(t *testing.T) {
		period := 1
		limit := 2

		mockRepoRental.EXPECT().GetTopAuthors(period, limit).Return([]dao2.TopAuthor{}, errors.New("Error"))

		_, err := libraryUseCase.GetTopAuthors(period, limit)

		assert.Error(t, err)
	})

	//16. GetTopAuthors success
	t.Run("GetTopAuthors success", func(t *testing.T) {
		period := 1
		limit := 2

		mockRepoRental.EXPECT().GetTopAuthors(period, limit).Return([]dao2.TopAuthor{}, nil)

		_, err := libraryUseCase.GetTopAuthors(period, limit)

		assert.NoError(t, err)
	})
}
