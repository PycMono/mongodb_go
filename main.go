package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"moqikaka.com/mongodb_go/src/dal"
	"time"
)

func main() {
	dal.StartConn()
	db := dal.GetDB()
	fmt.Println(db.Name())
	conllection := db.Collection("p_copy")

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, _ := conllection.Find(ctx, bson.D{})

	temp := new(Result)
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		fmt.Println(cur.Current)
		err := cur.Decode(&temp)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(temp.Name)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(11)
}

type Result struct {
	Name   string
	Hobbys string
}
