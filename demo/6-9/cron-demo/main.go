package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func main() {

	//设置时区
	//通过加载环境变量或者操作操作系统或者go内置的包或文件中加载指定时区
	//loc, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil {
	//	fmt.Println("load location error:", err)
	//	return
	//}

	// 创建东八区时区
	loc := time.FixedZone("CST", 8*3600)

	fmt.Println(time.Now().In(loc))
	fmt.Println(time.Now().In(time.UTC))

	// 创建一个cron对象，WithLocation设置时区，WithSeconds秒级别支持
	c := cron.New(cron.WithLocation(loc), cron.WithSeconds())

	//注意：cron表达式占位符可以有5个也可以有6个，取决于在cron.New时是否使用cron.WithSeconds()配置
	//当使用cron.WithSeconds()配置时，cron表达式占位符有5个，   从左到右依次表示分，时，日，月，周
	//当不使用cron.WithSeconds()配置时，cron表达式占位符有6个，从左到右依次表示秒，分，时，日，月，周

	//当cron表达式占位符为*号则表示每隔这个时间就会执行一次
	//比如：
	//	  * * * * *：每分钟执行一次任务。
	//	  0 * * * *：每小时执行一次任务。
	//	  0 0 * * *：每天午夜执行一次任务。
	//	  0 0 * * 1：每周一午夜执行一次任务。
	//	  0 0 1 * *：每月第一天午夜执行一次任务。

	//逗号和斜杠可以用来表示多个数值和时间间隔
	//	 0 30 1 ? * 1-5：表示周一到周五晚上一点半执行
	//  0 0/5 14 * * ? 表示在每天下午2点到下午2:55期间的每5分钟执行一次
	// 0 0-5 14 * * ? 表示在每天下午2点到下午2:05期间的每1分钟执行一次
	// 0 15,30,45 * * * ? 表示在每小时的15分、30分和45分执行一次

	// 添加一个每分钟执行一次的任务

	id1, err := c.AddFunc("* * * * *", func() {
		fmt.Println(time.Now().Format(time.DateTime), "id1 run every minutes")
	})
	if err != nil {
		log.Println("id1", err)
	}
	// 添加一个每秒钟执行一次的任务
	id2, err := c.AddFunc("* * * * * *", func() {
		fmt.Println(time.Now().Format(time.DateTime), "id2 run every seconds")
	})
	if err != nil {
		log.Println("id2", err)
	}

	job := myJob{}
	job.name = "id3"
	// 添加一个每小时执行一次的任务
	id3, err := c.AddJob("0 * * * *", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(job))
	if err != nil {
		log.Println("id3", err)
	}

	//这个规则看上去简单，在实际使用过程中却很容易踩坑
	//下面这个任务是每天凌晨2点执行吗？
	job.name = "id4"
	_, err = c.AddJob("* * 6 * * *", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(job))
	if err != nil {
		log.Println("id4", err)
	}
	//其实并不是，它表示从2点开始，每秒都执行，如果我们希望只在2点执行一次，则需要将秒和分钟设置成具体的时间
	job.name = "id5"
	_, err = c.AddJob("0 0 2 * * *", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(&myJob{}))
	if err != nil {
		log.Println("id5", err)
	}
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
type myJob struct {
	name string
}

// 实现Job接口的Run方法
func (j myJob) Run() {
	fmt.Println(time.Now().Format(time.DateTime), j.name)
}
