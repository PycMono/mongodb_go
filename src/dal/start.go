package dal

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	// connection object of mongodb
	db *mongo.Database
)

// connect mongodb
// parameter：nothing
// result：nothing
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

	// get db object by database name
	db = client.Database(getDatabaseName())
}

// get mongodb object
// parameter：nothing
// result：
// 1.database object
func GetDB() *mongo.Database {
	return db
}

// ret database name
// parameter：nothing
// result：
// database name
func getDatabaseName() string {
	return "pyc_test"
}
