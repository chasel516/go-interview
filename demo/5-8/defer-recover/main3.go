package main

func main() {
	doSomething1()
}

func doSomething1() (err error) {

	defer func() {
		err, _ = recover().(error)
	}()
	doStep_1()
	doStep_2()
	doStep_3()
	return
}

func doStep_1() {
	var err error
	var done bool
	// do something for doStep1
	if err != nil {
		panic(err)
	}
	if done {
		panic(nil)
	}
}

func doStep_2() {

	var err error
	var done bool
	// do something for doStep2
	if err != nil {
		panic(err)
	}
	if done {
		panic(nil)
	}
}

func doStep_3() {
	var err error
	var done bool
	// do something for doStep3
	if err != nil {
		panic(err)
	}
	if done {
		panic(nil)
	}
}
