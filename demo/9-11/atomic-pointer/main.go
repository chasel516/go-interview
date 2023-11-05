package main

import (
	"log"
	"sync/atomic"
	"unsafe"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type User struct {
	Name string
	Age  int
}

func main() {
	var ptr unsafe.Pointer

	// 存储一个指针
	user1 := &User{Name: "test1", Age: 30}
	log.Printf("user1.ptr:%p ; ptr.ptr:%v", user1, ptr)
	atomic.StorePointer(&ptr, unsafe.Pointer(user1))

	//ptr本身的就是指针，不需要使用%p进行格式化输出
	log.Printf("user1.ptr:%p ; ptr.ptr:%v", user1, ptr)

	// 加载指针并转换为对应类型
	user2 := (*User)(atomic.LoadPointer(&ptr))
	log.Printf("%+v", *user2)
	//读取出来的user2的地址跟ptr一致
	log.Printf("user2.ptr:%p ; ptr.ptr:%v", user2, ptr)

	// 交换指针
	user3 := &User{Name: "test2", Age: 25}
	//将ptr的地址指向user3
	oldValue := atomic.SwapPointer(&ptr, unsafe.Pointer(user3))
	log.Println("oldValue:", (*User)(oldValue), "newValue:", (*User)(atomic.LoadPointer(&ptr)))
	//ptr的地址发生了改变，指向了user3
	log.Printf("user3.ptr:%p ; ptr.ptr:%v", user3, ptr)

	// 比较并交换指针
	swapped := atomic.CompareAndSwapPointer(&ptr, unsafe.Pointer(user2), unsafe.Pointer(user1))
	log.Println(swapped)

	user4 := &User{Name: "test2", Age: 25}
	log.Printf("user3.ptr:%p ; user4.ptr:%p", user3, user4)
	swapped = atomic.CompareAndSwapPointer(&ptr, unsafe.Pointer(user4), unsafe.Pointer(user1))
	//这里user4的值虽然跟user3相同，但指向的地址不同，所以还是会替换失败
	log.Println(swapped)
	swapped = atomic.CompareAndSwapPointer(&ptr, unsafe.Pointer(user3), unsafe.Pointer(user1))
	log.Println(swapped)
	//将ptr的地址指向了user1
	log.Printf("user1.ptr:%p ; ptr.ptr:%v", user1, ptr)
	log.Printf("user1:%+v ; user3:%+v", user1, user3)

	// 加载交换后的指针并转换为对应类型
	user5 := (*User)(atomic.LoadPointer(&ptr))
	log.Println(user5)

}
