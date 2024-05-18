package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "93.125.114.79 - admin [19/Jun/2023:23:08:43 +0300] \"GET / HTTP/1.1\" 404 1133 \"-\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36\""
	parts := strings.Split(s, " ")
	var result []string
	inQuotes := false
	currentPart := ""
	for _, part := range parts {
		if strings.HasPrefix(part, "\"") && strings.HasSuffix(part, "\"") {
			inQuotes = false
			currentPart += part[1 : len(part)-1]
			result = append(result, currentPart)
			currentPart = ""
		} else if strings.HasPrefix(part, "\"") || strings.HasPrefix(part, "[") {
			inQuotes = true
			currentPart = part[1:]
		} else if strings.HasSuffix(part, "\"") || strings.HasSuffix(part, "]") {
			inQuotes = false
			currentPart += " " + part[:len(part)-1]
			result = append(result, currentPart)
			currentPart = ""
		} else if inQuotes {
			currentPart += " " + part
		} else {
			result = append(result, part)
		}
	}
	fmt.Println(result)
}
