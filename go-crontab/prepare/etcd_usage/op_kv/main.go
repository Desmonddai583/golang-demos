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

	// 创建Op: operation
	putOp := clientv3.OpPut("/cron/jobs/job8", "123123123")

	// 执行OP
	opResp, err := kv.Do(context.TODO(), putOp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入Revision:", opResp.Put().Header.Revision)

	// 创建Op
	getOp := clientv3.OpGet("/cron/jobs/job8")

	// 执行OP
	opResp, err = kv.Do(context.TODO(), getOp)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印
	fmt.Println("数据Revision:", opResp.Get().Kvs[0].ModRevision)
	fmt.Println("数据value:", string(opResp.Get().Kvs[0].Value))

}
