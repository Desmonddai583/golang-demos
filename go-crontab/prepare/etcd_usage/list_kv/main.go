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

	kv := clientv3.NewKV(client)

	kv.Put(context.TODO(), "/cron/jobs/job2", "{...}")

	// 读取/cron/jobs/为前缀的所有key
	getResp, err := kv.Get(context.TODO(), "/cron/jobs/", clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getResp.Kvs)
}
