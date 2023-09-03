package main

func main() {
	c := make(chan bool)
	close(c)
	select {

	//向关闭的channel写数据会导致panic
	case c <- true:
		//关闭的channel可以正常读取数据
	case <-c:
	}
}
