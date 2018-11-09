package main

import (
	"context"
	"fmt"

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

// FindByJobName jobName过滤条件
type FindByJobName struct {
	JobName string `bson:"jobName"`
}

func main() {
	// mongodb读取回来的是bson, 需要反序列为LogRecord对象

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

	// 按照jobName字段过滤, 想找出jobName=job10, 找出5条
	cond := &FindByJobName{JobName: "job10"} // {"jobName": "job10"}

	// 查询（过滤 +翻页参数）
	opts := options.Find()
	opts = opts.SetSkip(0) // 偏移量
	opts = opts.SetLimit(2)
	cursor, err := collection.Find(context.TODO(), cond, opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 延迟释放游标
	defer cursor.Close(context.TODO())

	// 6, 遍历结果集
	for cursor.Next(context.TODO()) {
		// 定义一个日志对象
		record := &LogRecord{}

		// 反序列化bson到对象
		err = cursor.Decode(record)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 把日志行打印出来
		fmt.Println(*record)
	}
}
