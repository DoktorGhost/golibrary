package entities

type BookTable struct {
	ID       int
	Title    string
	AuthorID int
}

type AuthorTable struct {
	ID   int
	Name string
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

type AuthorRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type BookRequest struct {
	Title    string `json:"title"`
	AuthorID int    `json:"authorID"`
}
