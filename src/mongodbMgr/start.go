package mongodbMgr

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	// 数据库对象
	db *mongo.Database
)

// 连接mongodb数据库
func StartConn() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(fmt.Sprintf("connect mongodb error ,error=%v", err))
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(fmt.Sprintf("Ping mongodb error ,error=%v", err))
		return
	}

	// 获取对应的数据库
	db = client.Database("pyc_test")
}

// 获取db数据库
// 参数：无
// 返回值：
// 1.数据库对象
func GetDB() *mongo.Database {
	return db
}
