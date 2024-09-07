package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Name string `example:"test"`
}

func main() {
	str := MyStruct{
		Name: "ttt",
	}

	sf := reflect.VisibleFields(reflect.TypeOf(str))

	for _, val := range sf {
		fmt.Println(val.Tag.Get("example"))
	}

}
