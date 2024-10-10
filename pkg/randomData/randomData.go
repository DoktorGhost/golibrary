package randomData

import (
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

func GenerateName() (string, string, string) {
	name := gofakeit.LastName()
	surname := gofakeit.LastName()
	patronymic := gofakeit.LastName()
	return name, surname, patronymic
}

func GenerateTitleBook() string {
	num := rand.Intn(4) + 1
	bookTitle := gofakeit.Sentence(num)
	return bookTitle
}
