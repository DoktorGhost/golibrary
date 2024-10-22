package entities

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterData struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type UserTable struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
	FullName     string `json:"full_name"`
}
