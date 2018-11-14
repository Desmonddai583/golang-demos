package main

import (
	"fmt"
	"golang-demos/go-crontab/crontab/master"
	"runtime"
)

func initArgs() {

}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)

	// 初始化线程
	initEnv()

	// 加载配置
	err = master.InitConfig("")
	if err != nil {
		fmt.Println(err)
	}

	// 启动API HTTP服务
	err = master.InitAPIServer()
	if err != nil {
		fmt.Println(err)
	}

	return
}
