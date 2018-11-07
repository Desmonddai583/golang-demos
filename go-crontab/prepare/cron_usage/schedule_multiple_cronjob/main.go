package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

// CronJob 代表一个任务
type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time // 通过expr.Next(now)获取
}

func main() {
	scheduleTable := make(map[string]*CronJob)

	now := time.Now()

	// 定义2个cronjob

	expr := cronexpr.MustParse("*/5 * * * * * *")
	cronJob := &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}

	// 任务注册到调度表
	scheduleTable["job1"] = cronJob

	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}

	scheduleTable["job2"] = cronJob

	// 需要有1个调度协程, 它定时检查所有的Cron任务, 谁过期了就执行谁
	go func() {
		// 定时检查一下任务调度表
		for {
			now := time.Now()

			for jobName, cronJob := range scheduleTable {
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					go func(jobName string) {
						fmt.Println("执行:", jobName)
					}(jobName)

					// 计算下一次调度时间
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println(jobName, "下次执行时间:", cronJob.nextTime)
				}
			}

			// 睡眠100毫秒
			select {
			case <-time.NewTimer(100 * time.Millisecond).C: // 将在100毫秒可读，返回
			}
		}
	}()

	// 防止程序直接退出
	time.Sleep(100 * time.Second)
}
