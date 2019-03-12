package main

import (
	"flag"
	"fmt"
	"golang-demos/go-crontab/crontab/master"
	"runtime"
	"time"
)

var (
	confFile string
)

func initArgs() {
	flag.StringVar(&confFile, "config", "./master.json", "指定master.json")
	flag.Parse()
}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)

	// 初始化命令行参数
	initArgs()

	// 初始化线程
	initEnv()

	// 加载配置
	if err = master.InitConfig(confFile); err != nil {
		fmt.Println(err)
	}

	// 初始化服务发现模块
	if err = master.InitWorkerMgr(); err != nil {
		fmt.Println(err)
	}

	// 日志管理器
	if err = master.InitLogMgr(); err != nil {
		fmt.Println(err)
	}

	// 任务管理器
	if err = master.InitJobMgr(); err != nil {
		fmt.Println(err)
	}

	// 启动API HTTP服务
	if err = master.InitAPIServer(); err != nil {
		fmt.Println(err)
	}

	// 正常退出
	for {
		time.Sleep(1 * time.Second)
	}
}
