package validator

import (
	"fmt"
	"strings"
	"unicode"
)

func validStr(name string) bool {
	for i, char := range name {
		// Проверяем, что символ является буквой или дефисом
		if !unicode.IsLetter(char) && char != '-' {
			return false
		}

		// Проверяем, что дефис не в начале или в конце строки
		if char == '-' && (i == 0 || i == len(name)-1) {
			return false
		}
	}
	return true
}

func Valid(name, surname, patronymic string) (string, error) {
	if !validStr(name) {
		return "", fmt.Errorf("имя не должно содержать цифры или символы")
	}
	if !validStr(surname) {
		return "", fmt.Errorf("фамилия не должна содержать цифры или символы")
	}
	if !validStr(patronymic) {
		return "", fmt.Errorf("отчество не должно содержать цифры или символы")
	}
	fullName := strings.TrimSpace(name + " " + surname + " " + patronymic)
	return fullName, nil
}
