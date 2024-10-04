package services

import (
	"fmt"

	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/repositories/postgres/dao"
)

type RentalsRepository interface {
	CreateRentalsInfo(userID, bookID int) (int, error)
	GetRentalsInfoByID(id int) (dao.RentalsTable, error)
	UpdateRentalsInfo(rentals dao.RentalsTable) error
	DeleteRentalsInfo(id int) error
	CreateRentals(bookID int) error
	GetRentalsByID(id int) (int, error)
	UpdateRentals(id, rentals_id int) error
	DeleteRentals(id int) error
	GetActiveRentals() (map[int][]int, error)
}

type RentalService struct {
	repo RentalsRepository
}

func NewRentalService(repo RentalsRepository) *RentalService {
	return &RentalService{repo: repo}
}

func (s *RentalService) CreateRentalInfo(userID, bookID int) (int, error) {
	rentalID, err := s.repo.CreateRentalsInfo(userID, bookID)
	if err != nil {
		return 0, fmt.Errorf("ошибка создания записи аренды: %v", err)
	}
	return rentalID, nil
}

func (s *RentalService) GetRentalsInfoByID(id int) (dao.RentalsTable, error) {
	rentals, err := s.repo.GetRentalsInfoByID(id)
	if err != nil {
		return dao.RentalsTable{}, fmt.Errorf("ошибка получени записи по ID=%d: %v", id, err)
	}
	return rentals, nil
}

func (s *RentalService) UpdateRentalsInfo(rentals dao.RentalsTable) error {
	err := s.repo.UpdateRentalsInfo(rentals)
	if err != nil {
		return fmt.Errorf("ошибка обновления записи аренды: %v", err)
	}
	return nil
}

func (s *RentalService) DeleteRentalsInfo(id int) error {
	err := s.repo.DeleteRentalsInfo(id)
	if err != nil {
		return fmt.Errorf("ошибка удаления записи аренды: %v", err)
	}
	return nil
}

func (s *RentalService) CreateRentals(bookID int) error {
	err := s.repo.CreateRentals(bookID)
	if err != nil {
		return fmt.Errorf("ошибка создания записи статуса книги: %v", err)
	}
	return nil
}

func (s *RentalService) GetRentalsByID(id int) (int, error) {
	rentalID, err := s.repo.GetRentalsByID(id)
	if err != nil {
		return 0, fmt.Errorf("ошибка получения статуса книги: %v", err)
	}
	return rentalID, nil
}

func (s *RentalService) UpdateRentals(id, rentalsID int) error {
	err := s.repo.UpdateRentals(id, rentalsID)
	if err != nil {
		return fmt.Errorf("ошибка обновления записи аренды: %v", err)
	}
	return nil
}

func (s *RentalService) DeleteRentals(id int) error {
	err := s.repo.DeleteRentals(id)
	if err != nil {
		return fmt.Errorf("ошибка удаления записи статуса книги: %v", err)
	}
	return nil
}

func (s *RentalService) GetActiveRentals() (map[int][]int, error) {
	rentalID, err := s.repo.GetActiveRentals()
	if err != nil {
		return nil, fmt.Errorf("ошибка получения записей активной аренды: %v", err)
	}
	return rentalID, nil
}
