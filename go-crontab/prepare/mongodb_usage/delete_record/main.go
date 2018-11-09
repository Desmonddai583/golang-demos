package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// TimeBeforeCond startTime小于某时间
// {"$lt": timestamp}
type TimeBeforeCond struct {
	Before int64 `bson:"$lt"`
}

// DeleteCond {"timePoint.startTime": {"$lt": timestamp} }
type DeleteCond struct {
	beforeCond TimeBeforeCond `bson:"timePoint.startTime"`
}

func main() {
	// 建立连接
	client, err := mongo.Connect(context.TODO(), "mongodb://192.168.10.130:27017")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 选择数据库my_db
	database := client.Database("cron")

	// 选择表my_collection
	collection := database.Collection("log")

	// 要删除开始时间早于当前时间的所有日志($lt是less than)
	// delete({"timePoint.startTime": {"$lt": 当前时间}})
	delCond := &DeleteCond{beforeCond: TimeBeforeCond{Before: time.Now().Unix()}}

	// 执行删除
	delResult, err := collection.DeleteMany(context.TODO(), delCond)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("删除的行数:", delResult.DeletedCount)
}
