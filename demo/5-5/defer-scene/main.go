package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

// 关闭文件句柄
func readFile(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // 使用 defer 语句关闭文件
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// 关闭http响应
func request() {
	url := "http://www.baidu.com"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close() // 关闭响应体

	// 处理响应数据
}

// 关闭数据库链接
func queryDatabase() error {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		return err
	}
	defer db.Close() // 在函数返回前关闭数据库连接
	// 查询数据库
	return nil
}

var mu sync.Mutex
var balance int

// 释放锁
func lockRelease(amount int) {
	mu.Lock()
	defer mu.Unlock() // 在函数返回前解锁
	balance += amount
}

// 捕获异常
func f() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(string(debug.Stack()))
		}
	}()
	panic("unknown")
}

// 取消任务
func runTask() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 在函数返回前取消任务
	go doTask(ctx)
	// 等待任务完成
	return nil
}

func doTask(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	}

	//do something
}

// 记录程序耗时
func trackTime() {
	start := time.Now()
	defer func() {
		log.Printf("Time took: %v", time.Since(start).Milliseconds())
	}()

	// 执行一些操作
	time.Sleep(time.Second * 3)
}
