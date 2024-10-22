package usecases

import (
	"errors"
	"fmt"
	services3 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/services"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/entities"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/repositories/postgres/dao"
	"time"

	services2 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/services"
	"github.com/DoktorGhost/golibrary/internal/core/user/services"
)

type LibraryUseCase struct {
	rentalService *services2.RentalService
	userService   *services.UserService
	bookService   *BookUseCase
	authorService *services3.AuthorService
}

func NewLibraryUseCase(
	rentalService *services2.RentalService,
	userService *services.UserService,
	bookService *BookUseCase,
	authorService *services3.AuthorService) *LibraryUseCase {
	return &LibraryUseCase{rentalService, userService, bookService, authorService}
}

// GiveBook выдать книгу
func (uc *LibraryUseCase) GiveBook(bookID, userID int) (int, error) {
	_, err := uc.userService.GetUserById(userID)
	if err != nil {
		return 0, fmt.Errorf("ошибка получения пользователя: %v", err)
	}

	rentalID, err := uc.rentalService.GetRentalsByID(bookID)
	if err != nil {
		return 0, fmt.Errorf("ошибка получения ID аренды: %v", err)
	}
	if rentalID > 0 {
		return 0, fmt.Errorf("книга уже выдана")
	}

	rentalID, err = uc.rentalService.CreateRentalInfo(userID, bookID)
	if err != nil {
		return 0, fmt.Errorf("ошибка создания записи аренды: %v", err)
	}

	err = uc.rentalService.UpdateRentals(bookID, rentalID)
	if err != nil {
		return 0, fmt.Errorf("ошибка обновления статуса книги: %v", err)
	}

	return rentalID, nil
}

// BackBook вернуть книгу
func (uc *LibraryUseCase) BackBook(bookID int) error {
	rentalID, err := uc.rentalService.GetRentalsByID(bookID)
	if err != nil {
		return fmt.Errorf("ошибка получения статуса книги: %v", err)
	}
	if rentalID == 0 {
		return errors.New("книга свободна")
	}
	err = uc.rentalService.UpdateRentals(bookID, 0)
	if err != nil {
		return fmt.Errorf("ошибка обновления статуса книги: %v", err)
	}

	rentalInfo, err := uc.rentalService.GetRentalsInfoByID(rentalID)
	if err != nil {
		return fmt.Errorf("ошибка получения данных аренды книги: %v", err)
	}
	rentalInfo.ReturnDate = time.Now()
	err = uc.rentalService.UpdateRentalsInfo(rentalInfo)
	if err != nil {
		return fmt.Errorf("ошибка обновления записи аренды: %v", err)
	}
	return nil
}

// GetUserRentals получить список пользователей с активной арендой
func (uc *LibraryUseCase) GetUserRentals() ([]entities.UserWithRentedBooks, error) {
	rentalsID, err := uc.rentalService.GetActiveRentals()
	if err != nil {
		return nil, fmt.Errorf("ошибка получения списка активной аренды: %v", err)
	}
	if len(rentalsID) < 1 {
		return nil, fmt.Errorf("записи не найдены: %v", err)
	}
	var result []entities.UserWithRentedBooks

	for userID, booksID := range rentalsID {
		var rental entities.UserWithRentedBooks

		username, err := uc.userService.GetUserById(userID)
		if err != nil {
			return nil, fmt.Errorf("ошибка получения автора: %v", err)
		}
		rental.ID = userID
		rental.Username = username

		for _, bookID := range booksID {
			book, err := uc.bookService.GetBookWithAuthor(bookID)
			if err != nil {
				return nil, err
			}
			rental.RentedBooks = append(rental.RentedBooks, book)
		}
		result = append(result, rental)
	}

	return result, nil
}

func (uc *LibraryUseCase) GetTopAuthors(period, limit int) ([]dao.TopAuthor, error) {
	authors, err := uc.rentalService.GetTopAuthorsByPeriod(period, limit)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения топ авторов за период %d: %v", period, err)
	}
	return authors, nil
}
