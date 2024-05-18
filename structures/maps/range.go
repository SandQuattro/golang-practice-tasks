package main

func main() {
	m := map[string]int{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}

	for b1, b2 := range m {
		println(b1, b2) // <- ?
	}

	a1, a2 := m["1"]
	println(a1, a2) // <- ?

}
