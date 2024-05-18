package main

import (
	"crypto/md5"
	"fmt"
	"testing"
)

var alphabet = []rune{'a', 'b', 'c', 'd', '1', '2', '3'}

// Генерация всех комбинаций паролей до заданной максимальной длины
func generateCombinations(alphabet []rune, length int) []string {
	var result []string

	var helper func(string, int)
	helper = func(combination string, maxLen int) {
		if len(combination) == maxLen {
			result = append(result, combination)
			return
		}
		for _, char := range alphabet {
			newCombination := combination + string(char)
			helper(newCombination, maxLen)
		}
	}

	for l := 1; l <= length; l++ {
		helper("", l)
	}
	return result
}

// Функция восстановления пароля
func RecoverPassword(hashedPassword []byte) string {
	maxLen := 4 // Установим максимальную длину пароля

	combinations := generateCombinations(alphabet, maxLen)
	for _, comb := range combinations {
		if md5.Sum([]byte(comb))[:] == hashedPassword {
			return comb
		}
	}
	return ""
}

func main() {
	// Пример использования
	password := "abc"
	recoveredPassword := RecoverPassword(hashPassword(password))
	fmt.Println("Recovered password:", recoveredPassword)
}

func TestRecoverPassword(t *testing.T) {
	for _, exp := range []string{
		"a",
		"12",
		"abc333d",
	} {
		t.Run(exp, func(t *testing.T) {
			act := RecoverPassword(hashPassword(exp))
			if act != exp {
				t.Error("recovered:", act, "expected:", exp)
			}
		})
	}
}

func hashPassword(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}
