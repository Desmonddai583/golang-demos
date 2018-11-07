package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	config := clientv3.Config{
		Endpoints:   []string{"192.168.10.130:4001"},
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 申请一个lease（租约）
	lease := clientv3.NewLease(client)

	// 申请一个10秒的租约
	leaseGrantResp, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 拿到租约的ID
	leaseID := leaseGrantResp.ID

	// 5秒后取消自动续租
	// ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)

	// 自动续租(每秒续一次)
	keepRespChan, err := lease.KeepAlive(context.TODO(), leaseID)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 处理续约应答的协程
	go func() {
		for {
			select {
			case keepResp := <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已经失效了")
					goto END
				} else { // 每秒会续租一次, 所以就会收到一次应答
					fmt.Println("收到自动续租应答:", keepResp.ID)
				}
			}
		}
	END:
	}()

	// 获得kv API子集
	kv := clientv3.NewKV(client)

	// Put一个KV, 让它与租约关联起来, 从而实现10秒后自动过期
	putResp, err := kv.Put(context.TODO(), "/cron/lock/job1", "", clientv3.WithLease(leaseID))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入成功:", putResp.Header.Revision)

	// 定时的看一下key过期了没有
	for {
		getResp, err := kv.Get(context.TODO(), "/cron/lock/job1")
		if err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("还没过期:", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}
}
