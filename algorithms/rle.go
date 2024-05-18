package main

import "fmt"

// необходимо реализовать кодирование строки алгоритмом rle, A8B6 и тд
func main() {
	in := "ААААААААBBBBBBAAAAACCCCDDEFF"
	fmt.Println(stringEncoding(in))
}

func stringEncoding(str string) string {
	out := ""
	var previous rune
	counter := 1
	for idx, r := range str {
		if idx == 0 {
			previous = r
			continue
		}
		if r != previous {
			out += fmt.Sprintf("%c%d", previous, counter)
			counter = 1
			previous = r
			continue
		}
		if r == previous {
			counter++
			continue
		}
	}
	out += fmt.Sprintf("%c%d", previous, counter)
	return out
}
