package main

func main() {

}
func doSomething() (err error) {
	isContinue, err := doStep1()
	if !isContinue {
		return err
	}
	isContinue, err = doStep2()
	if !isContinue {
		return err
	}
	isContinue, err = doStep3()
	if !isContinue {
		return err
	}
	return
}

func doStep1() (isContinue bool, err error) {
	// do something for doStep1
	return
}

func doStep2() (isContinue bool, err error) {
	// do something for doStep2
	return
}

func doStep3() (isContinue bool, err error) {
	// do something for doStep3
	return
}
