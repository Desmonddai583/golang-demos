package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// TimePoint 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// LogRecord 一条日志
// 注意字段需要大写(公有),否则序列化时不会被导出
type LogRecord struct {
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell命令
	Err       string    `bson:"err"`       // 脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点
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

	// 插入记录(bson)
	record := &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	result, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		fmt.Println(err)
		return
	}

	// _id: 默认生成一个全局唯一ID, ObjectID：12字节的二进制
	docID := result.InsertedID.(objectid.ObjectID)
	fmt.Println("自增ID:", docID.Hex())
}
