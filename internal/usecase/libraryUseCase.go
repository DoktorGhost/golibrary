package usecase

import (
	"errors"
	"github.com/DoktorGhost/golibrary/internal/services"
	"time"
)

type LibraryUseCase struct {
	rentalService services.RentalService
	userService   services.UserService
}

func NewLibraryUseCase(rentalService services.RentalService, userService services.UserService) *LibraryUseCase {
	return &LibraryUseCase{rentalService, userService}
}

// выдать книгу
func (uc *LibraryUseCase) GiveBook(bookID, userID int) (int, error) {
	_, err := uc.userService.GetUserById(userID)
	if err != nil {
		return -1, err
	}

	rentalID, err := uc.rentalService.GetRentalsByID(bookID)
	if err != nil {
		return -2, err
	}
	if rentalID > 0 {
		return -3, errors.New("книга уже выдана")
	}

	rentalID, err = uc.rentalService.CreateRentalInfo(userID, bookID)
	if err != nil {
		return -4, err
	}

	err = uc.rentalService.CreateRentals(bookID, rentalID)
	if err != nil {
		return -5, err
	}

	return rentalID, nil
}

// вернуть книгу
func (uc *LibraryUseCase) BackBook(bookID int) error {
	rentalID, err := uc.rentalService.GetRentalsByID(bookID)
	if err != nil {
		return err
	}
	if rentalID == 0 {
		return errors.New("книга не выдана")
	}
	err = uc.rentalService.UpdateRentals(bookID, 0)
	if err != nil {
		return err
	}

	rentalInfo, err := uc.rentalService.GetRentalsInfoByID(rentalID)
	if err != nil {
		return err
	}
	rentalInfo.ReturnDate = time.Now()
	err = uc.rentalService.UpdateRentalsInfo(rentalInfo)
	if err != nil {
		return err
	}
	return nil
}
