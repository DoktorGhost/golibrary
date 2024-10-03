package models

import "time"

type UserTable struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
	FullName     string `json:"full_name"`
}

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	RentedBooks Book   `json:"rented_books"`
}

type RegisterData struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type BookTable struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
}

type AuthorTable struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	ID     int         `json:"id"`
	Title  string      `json:"title"`
	Author AuthorTable `json:"author"`
}

type Author struct {
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Books []BookTable `json:"books"`
}

type RentalsTable struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	BookID     int       `json:"book_id"`
	RentalDate time.Time `json:"rental_date"`
	ReturnDate time.Time `json:"return_date"`
}
