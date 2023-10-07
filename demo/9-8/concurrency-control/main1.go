//缓冲通道可以被用做计数信号量 � 。 计数信号量可以被视为多主锁。如果一个缓冲通道的容量为N，那么它可以被看作是一个在任何时刻最多可有N个主人的锁。 上面提到的二元信号量是特殊的计数信号量，每个二元信号量在任一时刻最多只能有一个主人。
//计数信号量经常被使用于限制最大并发数。
//和将通道用做互斥锁一样，也有两种方式用来获取一个用做计数信号量的通道的一份所有权。1. 通过发送操作来获取所有权，通过接收操作来释放所有权；
//2. 通过接收操作来获取所有权，通过发送操作来释放所有权。下面是一个通过接收操作来获取所有权的例子：
 package main

 import (
 "log"
"time"
"math/rand"
)

type Seat int
 type Bar chan Seat

 func (bar Bar) ServeCustomer(c int) {
 log.Print("顾客#", c, "进入酒吧")
 seat := <- bar // 需要一个位子来喝酒
 log.Print("++ customer#", c, " drinks at seat#", seat)
 log.Print("++ 顾客#", c, "在第", seat, "个座位开始饮酒")
 time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
 log.Print("-- 顾客#", c, "离开了第", seat, "个座位")
 bar <- seat // 释放座位，离开酒吧20| }21|
 func main() {
 rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要24|
 bar24x7 := make(Bar, 10) // 此酒吧有10个座位26| // 摆放10个座位。
 for seatId := 0; seatId < cap(bar24x7); seatId++ {
 bar24x7 <- Seat(seatId) // 均不会阻塞29| }30|
 for customerId := 0; ; customerId++ {
 time.Sleep(time.Second)
 go bar24x7.ServeCustomer(customerId)
 }


 for {time.Sleep(time.Second)} // 睡眠不属于阻塞状态36| }
//在上例中，只有获得一个座位的顾客才能开始饮酒。 所以在任一时刻同时在喝酒的顾客数不会超过座位数10。
//上例main函数中的最后一行for循环是为了防止程序退出。 后面将介绍一种更好的实现此目的的方法。
//在上例中，尽管在任一时刻同时在喝酒的顾客数不会超过座位数10，但是在某一时刻可能有多于10个顾客进入了酒吧，因为某些顾客在排队等位子。 在上例中，每个顾客对应着一个协程。虽然协程的开销比系统线程小得多，但是如果协程的数量很多，则它们的总体开销还是不能忽略不计的。 所以，最好当有空位的时候才创建顾客协程。
//... // 省略了和上例相同的代码2|
func (bar Bar) ServeCustomerAtSeat(c int, seat Seat){
	 log.Print("++ 顾客#", c, "在第",
 } seat, "个座位开始饮酒")
	  time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
log.Print("-- 顾客#", c, "离开了第", seat, "个座位")
bar <- seat // 释放座位，离开酒吧8|
 }
 func main() {
 rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要12|
 bar24x7 := make(Bar, 10)
 for seatId := 0; seatId < cap(bar24x7); seatId++ {
 bar24x7 <- Seat(seatId)
 }

 // 这个for循环和上例不一样。
 for customerId := 0; ; customerId++ {
 time.Sleep(time.Second)
 seat := <- bar24x7 // 需要一个空位招待顾客
 go bar24x7.ServeCustomerAtSeat(customerId, seat)
 }
 for {time.Sleep(time.Second)}
 }
//在上面这个修改后的例子中，在任一时刻最多只有10个顾客协程在运行（但是在程序的生命期内，仍旧会有大量的顾客协程不断被创建和销毁）。
//在下面这个更加高效的实现中，在程序的生命期内最多只会有10个顾客协程被创建出来。
//第37章：通道用例大全

//... // 省略了和上例相同的代码2|
func (bar Bar) ServeCustomerAtSeat(consumers chan int) {
for c := range consumers {
seatId := <- bar
log.Print("++ 顾客#", c, "在第", seatId, "个座位开始饮酒")
time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
log.Print("-- 顾客#", c, "离开了第", seatId, "个座位")
bar <- seatId // 释放座位，离开酒吧10| }11| }12|
 func main() {
 rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要15|
 bar24x7 := make(Bar, 10)
 for seatId := 0; seatId < cap(bar24x7); seatId++ {
 bar24x7 <- Seat(seatId)
 }

 consumers := make(chan int)
 for i := 0; i < cap(bar24x7); i++ {
 go bar24x7.ServeCustomerAtSeat(consumers)
 }

 for customerId := 0; ; customerId++ {
 time.Sleep(time.Second)28| consumers <- customerId29| }
 }
//题外话：当然，如果我们并不关心座位号（这种情况在编程实践中很常见），则实际上bar24x7计数信号量是完全不需要的：
//... // 省略了和上例相同的代码2|
func ServeCustomer(consumers chan int) {
for c := range consumers {
log.Print("++ 顾客#", c, "开始在酒吧饮酒")
time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
log.Print("-- 顾客#", c, "离开了酒吧")8| }9| }10|
 func main() {
 rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要


 const BarSeatCount = 10
 consumers := make(chan int)
 for i := 0; i < BarSeatCount; i++ {
 go ServeCustomer(consumers)
 }

 for customerId := 0; ; customerId++ {
 time.Sleep(time.Second)22| consumers <- customerId23| }
 }
//通过发送操作来获取所有权的实现相对简单一些，省去了摆放座位的步骤。
package main

import (
"log"
"time"
"math/rand"
)

type Customer struct{id int}
 type Bar chan Customer

 func (bar Bar) ServeCustomer(c Customer) {
 log.Print("++ 顾客#", c.id, "开始饮酒")
 time.Sleep(time.Second * time.Duration(3 + rand.Intn(16)))
 log.Print("-- 顾客#", c.id, "离开酒吧")
 <- bar // 离开酒吧，腾出位子17| }18|
 func main() {
 rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要21|
 bar24x7 := make(Bar, 10) // 最对同时服务10位顾客23| for customerId := 0; ; customerId++ {24| time.Sleep(time.Second * 2)
 customer := Customer{customerId}26| bar24x7 <- customer // 等待进入酒吧27| go bar24x7.ServeCustomer(customer)28| }
 for {time.Sleep(time.Second)}
 }



//峰值限制（peak/burst limiting）
//将通道用做计数信号量用例和通道尝试（发送或者接收）操作结合起来可用实现峰值限制。 峰值限制的目的是防止过大的并发请求数。
//下面是对将通道用做计数信号量一节中的最后一个例子的简单修改，从而使得顾客不再等待而是离去或者寻找其它酒吧。
...
bar24x7 := make(Bar, 10) // 此酒吧只能同时招待10个顾客3| for customerId := 0; ; customerId++ {4| time.Sleep(time.Second)
consumer := Consumer{customerId}
select {
case bar24x7 <- consumer: // 试图进入此酒吧8| go bar24x7.ServeConsumer(consumer)9| default:
 log.Print("顾客#", customerId, "不愿等待而离去")
}
 }