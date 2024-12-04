package main

import "fmt"

func isMatch(string string, pattern string) bool {
	if len(pattern) != len(string) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if string[i] != pattern[i] && pattern[i] != '?' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isMatch("", ""))               // true
	fmt.Println(isMatch("1", "1"))             // true
	fmt.Println(isMatch("1010111", "101?11?")) //true

	fmt.Println(isMatch("", "101?11?"))        // false
	fmt.Println(isMatch("0010101", "101?11?")) // false

	// Тесты с полностью совпадающими строками
	fmt.Println(isMatch("12345", "12345")) // true
	fmt.Println(isMatch("00000", "00000")) // true

	// Тесты с использованием подстановочных символов
	fmt.Println(isMatch("abcde", "a?c?e")) // true
	fmt.Println(isMatch("11111", "1?1?1")) // true
	fmt.Println(isMatch("54321", "5?3?1")) // true

	// Тесты с несовпадающими строками
	fmt.Println(isMatch("12345", "54321")) // false
	fmt.Println(isMatch("abcde", "12345")) // false

	// Тесты с предельными случаями
	fmt.Println(isMatch("?", "?")) // true
	fmt.Println(isMatch("?", "1")) // false

	// Тесты с частичным использованием подстановочных символов
	fmt.Println(isMatch("hello", "h?ll?")) // true
	fmt.Println(isMatch("world", "w?r?d")) // true

	// Тесты с несовпадающей длиной
	fmt.Println(isMatch("short", "longer")) // false
	fmt.Println(isMatch("a", "abc"))        // false
}
