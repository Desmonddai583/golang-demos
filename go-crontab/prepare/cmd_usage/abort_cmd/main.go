package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {
	// 强制结束任务

	// 创建了一个结果队列
	resultChan := make(chan *result, 1000)

	// context: chan byte
	// cancelFunc: close(chan byte)
	ctx, cancelFunc := context.WithCancel(context.TODO())

	go func() {
		// 这里内部会调用一个select去监听ctx.Done()(会返回一个chan), select { case <- ctx.Done(); }
		// 如果监听到关闭就会kill掉当前子进程pid
		cmd := exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2; echo hello;")

		output, err := cmd.CombinedOutput()

		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()

	time.Sleep(1 * time.Second)

	// 取消上下文
	cancelFunc()

	// 在main协程里, 等待子协程的退出，并打印任务执行结果
	res := <-resultChan

	fmt.Println(res.err, string(res.output))
}
