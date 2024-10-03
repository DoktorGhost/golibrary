package randomData

import (
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
)

func GenerateName() (string, string, string) {
	name := gofakeit.LastName()
	lastName := gofakeit.LastName()
	patronymic := gofakeit.LastName()
	return name, lastName, patronymic
}

func GenerateTitleBook() string {
	num := rand.Intn(4) + 1
	bookTitle := gofakeit.Sentence(num)
	return bookTitle
}
