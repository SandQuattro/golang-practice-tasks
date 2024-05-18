package main

type Sum interface {
	Add(Sum) Sum
}

type IntSum struct {
	value int64
}

func main() {
	var a Sum
	a = &IntSum{value: 1}
	a = a.Add(&IntSum{value: 2})
	a = a.Add(&IntSum{value: 3})
}

func (i *IntSum) Add(a Sum) Sum {
	switch a.(type) {
	case *IntSum:
		return &IntSum{
			value: i.value + a.(*IntSum).value,
		}

	default:
		return nil
	}
}
