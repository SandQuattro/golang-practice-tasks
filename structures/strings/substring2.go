package main

import (
	"fmt"
	"strings"
)

func main() {
	str := `Fetching origin
	From /home/git/smart-analytics-core
	c716fef..c6dccf4  main       -> origin/main
`
	sub := str[strings.Index(str, "-> origin")+10:]
	fmt.Print(sub[:strings.Index(sub, "\n")])
}
