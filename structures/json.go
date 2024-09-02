package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	X int    `json:"x"`
	y string `json:"y,omitempty"`
}

func testData() {
	in := Data{1, "two"}
	fmt.Printf("%#v\n", in) // ?

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded)) // ?

	var out Data
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out) // ?
}

func main() {
	testData()
}
