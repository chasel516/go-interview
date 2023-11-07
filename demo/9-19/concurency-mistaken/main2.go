package main

func main() {
	var sl []int
	//var ok bool
	var ch = make(chan struct{}, 1)
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	go func() {
		sl = make([]int, 1)
		//wg.Done()
		//ok = true
		ch <- struct{}{}

	}()

	//for !ok {
	//	time.Sleep(time.Millisecond)
	//}
	<-ch
	//wg.Wait()
	sl[0] = 1
}
