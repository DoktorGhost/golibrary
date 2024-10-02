package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	FIO      string `json:"fio"`
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
