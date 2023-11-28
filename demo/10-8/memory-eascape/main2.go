package main

func main() {}
func f1() *int {
	var n1 int
	n1 = 42
	return &n1 //n1的地址被返回，导致n1逃逸到堆上
}
func f2() func() {
	var n2 int
	return func() { //闭包函数func发生逃逸
		n2 += 1 //n2发生逃逸
		n21 := 1
		n21 += 1
	}
}

func f3() func() int {
	n3 := 1
	return func() int {
		n31 := 1
		n31 += 1
		return n3 // 函数返回后n3还在使用，导致n3逃逸到堆上
	}
}

func f4() {
	var n4 int
	var ch = make(chan int)
	ch <- 1
	go func() { //闭包函数func发生逃逸
		n4 += 1 //n4 发生逃逸
		<-ch    //通道类型并没发生逃逸
		n41 := 1
		n41 += 1
	}()
}
