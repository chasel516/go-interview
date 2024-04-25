package main

func main() {
	select {}
}

//对于空的select语句，程序会被阻塞，准确的说是当前协程被阻塞，同时golang自带死锁检测机制，当发现当前协程再也没机会被唤醒时，则会panic
