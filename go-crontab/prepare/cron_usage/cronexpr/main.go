package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

func main() {
	// cronexpr支持的力度更细,有秒和年(2018-2099)
	// 还有一个MustParse方法,如果出错直接报错
	expr, err := cronexpr.Parse("*/5 * * * * * *")
	if err != nil {
		fmt.Println(err)
		return
	}

	now := time.Now()

	nextTime := expr.Next(now)

	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了:", nextTime)
	})

	// 防止程序直接退出
	time.Sleep(5 * time.Second)
}
