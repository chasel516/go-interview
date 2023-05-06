package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {

	//设置时区
	//loc, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil {
	//	fmt.Println("load location error:", err)
	//	return
	//}

	// 创建东八区时区
	loc := time.FixedZone("CST", 8*3600)

	fmt.Println(time.Now().In(loc))
	fmt.Println(time.Now().In(time.UTC))

	// 创建一个cron对象
	c := cron.New(cron.WithLocation(loc))

	// 添加一个每分钟执行一次的任务
	id1, _ := c.AddFunc("* * * * *", func() {
		fmt.Println("run every minutes")
	})

	// 添加一个每分钟执行一次的任务
	id2, _ := c.AddFunc("* * * * * *", func() {
		fmt.Println("run every seconds")
	})

	// 添加一个每小时执行一次的任务
	id3, _ := c.AddJob("0 * * * *", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(&myJob{}))

	// 启动定时任务
	c.Start()

	// 查看所有的定时任务
	for _, e := range c.Entries() {
		fmt.Printf("ID: %d, Next: %v, Prev: %v, Schedule: %v\n", e.ID, e.Next, e.Prev, e.Schedule)
	}

	// 查看指定的定时任务
	e := c.Entry(id1)
	fmt.Printf("ID: %d, Next: %v, Prev: %v, Schedule: %v\n", e.ID, e.Next, e.Prev, e.Schedule)

	e = c.Entry(id2)
	fmt.Printf("ID: %d, Next: %v, Prev: %v, Schedule: %v\n", e.ID, e.Next, e.Prev, e.Schedule)

	// 删除指定的定时任务
	c.Remove(id3)

	// 停止定时任务
	defer c.Stop()

	// 主程序阻塞
	select {}
}

// 定义一个实现了Job接口的结构体
type myJob struct{}

// 实现Job接口的Run方法
func (j myJob) Run() {
	fmt.Println("Do something")
}
