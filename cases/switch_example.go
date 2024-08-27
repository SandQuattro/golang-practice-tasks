package main

import (
	"flag"
	"log"
)

func main() {
	var s string
	flag.StringVar(&s, "s", "", "String")
	flag.Parse()

	switch s {
	case "1", "2", "3":
		log.Println("case 1")
	case "4", "5", "6":
		log.Println("case 2")
	default:
		log.Println("default")
	}
}
