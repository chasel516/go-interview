package main

func main() {
	ch := make(chan int)
	//<-ch
	//go fmt.Println(<-ch)

	//go func() {
	//	fmt.Println(<-ch)
	//}()
	//ch <- 1

	close(ch)
	ch <- 1
	//time.Sleep(time.Second)
}
