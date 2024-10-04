package services

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/models"
	"github.com/DoktorGhost/golibrary/internal/services/crud"
)

type RentalService struct {
	repo crud.RentalsRepository
}

func NewRentalService(repo crud.RentalsRepository) *RentalService {
	return &RentalService{repo: repo}
}

func (s *RentalService) CreateRentalInfo(userID, bookID int) (int, error) {
	rentalID, err := s.repo.CreateRentalsInfo(userID, bookID)
	if err != nil {
		return 0, fmt.Errorf("ошибка создания записи: %v", err)
	}
	return rentalID, nil
}

func (s *RentalService) GetRentalsInfoByID(id int) (models.RentalsTable, error) {
	rentals, err := s.repo.GetRentalsInfoByID(id)
	if err != nil {
		return models.RentalsTable{}, err
	}
	return rentals, nil
}

func (s *RentalService) UpdateRentalsInfo(rentals models.RentalsTable) error {
	err := s.repo.UpdateRentalsInfo(rentals)
	if err != nil {
		return err
	}
	return nil
}

func (s *RentalService) DeleteRentalsInfo(id int) error {
	err := s.repo.DeleteRentalsInfo(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *RentalService) CreateRentals(bookID int) error {
	err := s.repo.CreateRentals(bookID)
	if err != nil {
		return fmt.Errorf("ошибка создания записи: %v", err)
	}
	return nil
}

func (s *RentalService) GetRentalsByID(id int) (int, error) {
	rentalID, err := s.repo.GetRentalsByID(id)
	if err != nil {
		return 0, err
	}
	return rentalID, nil
}

func (s *RentalService) UpdateRentals(id, rentals_id int) error {
	err := s.repo.UpdateRentals(id, rentals_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *RentalService) DeleteRentals(id int) error {
	err := s.repo.DeleteRentals(id)
	if err != nil {
		return err
	}
	return nil
}
