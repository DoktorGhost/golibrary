package dao

import "time"

type RentalsTable struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	BookID     int       `json:"book_id"`
	RentalDate time.Time `json:"rental_date"`
	ReturnDate time.Time `json:"return_date"`
}

type TopAuthor struct {
	Name      string `json:"name"`
	CountRent int    `json:"count_rent"`
}
