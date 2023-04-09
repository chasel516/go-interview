package main

import (
	"context"
	"time"
)

func main() {
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, time.Second*5)
	defer cancel()
	// 在函数中检查ctx.Done()来判断是否需要取消
	if err := longRunningFunction(ctx, "param1", 123); err != nil {
		// 如果发生错误，调用cancel()来取消请求
		cancel()
	}
}

func longRunningFunction(ctx context.Context, arg1 string, arg2 int) error {
	select {
	case <-time.After(time.Second * 5):
		// 执行操作
	case <-ctx.Done():
		// 返回错误,取消请求
		return ctx.Err()
	}
	return nil
}
