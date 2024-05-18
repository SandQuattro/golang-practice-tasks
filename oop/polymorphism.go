package main

func specificMethod1() error {
	println("spec method 1")
	return nil
}

func specificMethod2() error {
	println("spec method 2")
	return nil
}

func main() {
	err := universalMethod(specificMethod1)
	if err != nil {
		return
	}

	err = universalMethod(specificMethod2)
	if err != nil {
		return
	}
}

func universalMethod(f func() error) error {
	err := f()
	if err != nil {
		return err
	}
	return nil
}
