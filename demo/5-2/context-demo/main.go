package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// 模拟一个耗时的操作，比如数据库查询或者网络请求
func slowOperation(ctx context.Context) (string, error) {
	//获取request_id
	fmt.Println("get request_id from ctx", ctx.Value("request_id"))

	// 使用一个 select 语句来监听 context 的状态变化
	select {
	case <-time.After(3 * time.Second): // 模拟操作需要 3 秒钟才能完成
		return "Done", nil
	case <-ctx.Done(): // 如果 context 被取消或者超时，返回相应的错误信息
		return "", ctx.Err()
	}
}

// 模拟一个 HTTP 服务器的处理函数，使用 context 来控制超时
func handler(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()
	fmt.Println("gen request_id ", requestID)
	//传递request_id
	ctr := context.WithValue(r.Context(), "request_id", requestID)

	// 从请求中获取 context，并设置一个 1 秒钟的超时时间
	ctx, cancel := context.WithTimeout(ctr, 1*time.Second)

	//另一种超时设置方式
	//deadline := time.Now().Add(1 * time.Second)
	//ctx, cancel := context.WithDeadline(ctr, deadline)

	// 在函数返回时调用取消函数
	defer cancel()

	// 调用耗时的操作，并传递 context
	result, err := slowOperation(ctx)
	if err != nil {
		// 如果操作失败，返回错误信息和状态码
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
		return
	}
	// 如果操作成功，返回结果
	fmt.Fprintln(w, result)
}

func main() {
	// 创建一个 HTTP 服务器，并注册处理函数
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost:9090")
	// 启动服务器，并监听 8080 端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}

}
