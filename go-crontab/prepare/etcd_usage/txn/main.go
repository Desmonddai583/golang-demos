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

	// 1. 上锁 (创建租约, 自动续租, 拿着租约去抢占一个key)
	lease := clientv3.NewLease(client)

	// 申请一个5秒的租约
	leaseGrantResp, err := lease.Grant(context.TODO(), 5)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 拿到租约的ID
	leaseID := leaseGrantResp.ID

	// 准备一个用于取消自动续租的context
	ctx, cancelFunc := context.WithCancel(context.TODO())

	// 确保函数退出后, 自动续租会停止
	defer cancelFunc()
	defer lease.Revoke(context.TODO(), leaseID)

	// 5秒后会取消自动续租
	keepRespChan, err := lease.KeepAlive(ctx, leaseID)
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

	kv := clientv3.NewKV(client)

	// 创建事务
	txn := kv.Txn(context.TODO())

	// 定义事务
	// if不存在key, then设置它, else抢锁失败
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job9"), "=", 0)).
		Then(clientv3.OpPut("/cron/lock/job9", "xxx", clientv3.WithLease(leaseID))).
		Else(clientv3.OpGet("/cron/lock/job9"))

	// 提交事务
	txnResp, err := txn.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 判断是否抢到了锁
	if !txnResp.Succeeded {
		fmt.Println("锁被占用:", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	// 2, 处理业务

	fmt.Println("处理任务")
	time.Sleep(5 * time.Second)

	// 3, 释放锁(取消自动续租, 释放租约)
	// defer 会把租约释放掉, 关联的KV就被删除了
}
