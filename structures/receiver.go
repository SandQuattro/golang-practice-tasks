package main

type Receiver struct {
	text string
}

// func (r *Receiver) Error() string {
func (r Receiver) Error() string {
	return r.text
}

func f() error {
	return &Receiver{
		"hello",
	}
}

func main() {
	_ = f()
}
