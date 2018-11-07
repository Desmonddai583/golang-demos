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

	// 用于读写etcd的键值对
	kv := clientv3.NewKV(client)

	// 默认也会返回count,WithCountOnly是resp只返回count
	getResp, err := kv.Get(context.TODO(), "/cron/jobs/job1", clientv3.WithCountOnly())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getResp.Kvs, getResp.Count)
}
