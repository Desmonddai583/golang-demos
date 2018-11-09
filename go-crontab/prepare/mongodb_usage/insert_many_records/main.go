package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/options"
)

// TimePoint 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// LogRecord 一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell命令
	Err       string    `bson:"err"`       // 脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点
}

func main() {
	// 建立连接
	opts := options.Client().SetConnectTimeout(5 * time.Second)
	client, err := mongo.Connect(context.TODO(), "mongodb://192.168.10.130:27017", opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 选择数据库my_db
	database := client.Database("cron")

	// 选择表my_collection
	collection := database.Collection("log")

	// 插入记录(bson)
	record := &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	// 批量插入多条document
	logArr := []interface{}{record, record, record}

	result, err := collection.InsertMany(context.TODO(), logArr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// mongodb生成ID的算法
	// snowflake: 毫秒/微秒的当前时间 + 机器的ID + 当前毫秒/微秒内的自增ID(每当毫秒变化了, 会重置成0，继续自增)
	for _, insertID := range result.InsertedIDs {
		// 拿着interface{}， 反射成objectID
		docID := insertID.(objectid.ObjectID)
		fmt.Println("自增ID:", docID.Hex())
	}
}
