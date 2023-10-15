package main

import (
	"log"
	"sync/atomic"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	var x int64
	//Store设置值；第一个参数传整型的地址，第二个参数为需要设置的值
	//可以确保在多个Goroutine在多核或者多个处理器有CPU Cache的情况下的原子操作
	//注：atomic提供的全部方法都能保证操作的原子性
	atomic.StoreInt64(&x, 1)
	log.Println(x)

	//Add增加值；第一个参数传整型的地址，第二个参数可以传负数
	atomic.AddInt64(&x, -1)

	//Load读取值；
	log.Println(atomic.LoadInt64(&x))

	//Swap强制交换,将第一个参数的值与第二个参数的值进行交换,并返回交换前的值
	//Swap的效果跟Store很像，但Swap会返回旧值
	oldValue := atomic.SwapInt64(&x, 2)
	log.Println(oldValue, x)

	//CompareAndSwap比较并交换，如果当前addr指向的值与给定的值相等才进行交换
	//下面代码表示的如果x等于3就将x的值设置为4.
	//由于此时addr指向的值为2，不等于3，所以下面的代码并不会将x设置为4
	ok := atomic.CompareAndSwapInt64(&x, 3, 4)
	log.Println(ok, x)

	ok = atomic.CompareAndSwapInt64(&x, 2, 3)
	log.Println(ok, x)

	//bool类型
	var b atomic.Bool
	b.Store(true)
	b.Swap(false)
	ok = b.CompareAndSwap(false, true)
	log.Println(ok, b)

	//Value类型
	var v atomic.Value
	type User struct {
		name string
		age  uint16
	}
	u := User{
		name: "test",
		age:  10,
	}
	v.Store(&u)
	u1 := User{
		name: "test1",
		age:  11,
	}
	v.Swap(&u1)
	//注意：两个参数都需要传地址
	v.CompareAndSwap(&u1, &u)
	log.Println(v.Load())
}
