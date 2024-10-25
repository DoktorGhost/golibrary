package services

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/repositories/postgres/dao"
	"github.com/DoktorGhost/golibrary/internal/metrics"
	"time"
)

//go:generate mockgen -source=$GOFILE -destination=./mock_rental.go -package=${GOPACKAGE}
type RentalsRepository interface {
	CreateRentalsInfo(userID, bookID int) (int, error)
	GetRentalsInfoByID(id int) (dao.RentalsTable, error)
	UpdateRentalsInfo(rentals dao.RentalsTable) error
	CreateRentals(bookID int) error
	GetRentalsByID(id int) (int, error)
	UpdateRentals(id, rentals_id int) error
	GetActiveRentals() (map[int][]int, error)
}

type RentalService struct {
	repo RentalsRepository
}

func NewRentalService(repo RentalsRepository) *RentalService {
	return &RentalService{repo: repo}
}

func (s *RentalService) CreateRentalInfo(userID, bookID int) (int, error) {
	start := time.Now()

	rentalID, err := s.repo.CreateRentalsInfo(userID, bookID)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("CreateRentalInfo", duration)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания записи аренды: %v", err)
	}

	return rentalID, nil
}

func (s *RentalService) GetRentalsInfoByID(id int) (dao.RentalsTable, error) {
	start := time.Now()

	rentals, err := s.repo.GetRentalsInfoByID(id)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("GetRentalsInfoByID", duration)

	if err != nil {
		return dao.RentalsTable{}, fmt.Errorf("ошибка получени записи по ID=%d: %v", id, err)
	}

	return rentals, nil
}

func (s *RentalService) UpdateRentalsInfo(rentals dao.RentalsTable) error {
	start := time.Now()

	err := s.repo.UpdateRentalsInfo(rentals)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("UpdateRentalsInfo", duration)

	if err != nil {
		return fmt.Errorf("ошибка обновления записи аренды: %v", err)
	}

	return nil
}

func (s *RentalService) CreateRentals(bookID int) error {
	start := time.Now()

	err := s.repo.CreateRentals(bookID)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("CreateRentals", duration)

	if err != nil {
		return fmt.Errorf("ошибка создания записи статуса книги: %v", err)
	}

	return nil
}

func (s *RentalService) GetRentalsByID(id int) (int, error) {
	start := time.Now()

	rentalID, err := s.repo.GetRentalsByID(id)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("AddAuthor", duration)

	if err != nil {
		return 0, fmt.Errorf("ошибка получения статуса книги: %v", err)
	}

	return rentalID, nil
}

func (s *RentalService) UpdateRentals(id, rentalsID int) error {
	start := time.Now()

	err := s.repo.UpdateRentals(id, rentalsID)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("UpdateRentals", duration)

	if err != nil {
		return fmt.Errorf("ошибка обновления записи аренды: %v", err)
	}

	return nil
}

func (s *RentalService) GetActiveRentals() (map[int][]int, error) {
	start := time.Now()

	rentalID, err := s.repo.GetActiveRentals()

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("GetActiveRentals", duration)

	if err != nil {
		return nil, fmt.Errorf("ошибка получения записей активной аренды: %v", err)
	}
	return rentalID, nil
}
