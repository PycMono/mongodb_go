package copydal

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"moqikaka.com/mongodb_go/src/dal"
	"moqikaka.com/mongodb_go/src/model/hero"
	"time"
)

// get all hero info by playerID
// parameter：
// playerID：unique id
// result：
// 1.hero list
func GetList(playerID string) []*hero.Hero {
	heroList := make([]*hero.Hero, 0, 10)

	db := dal.GetDB()
	conllection := db.Collection(getCollectionName())

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, _ := conllection.Find(ctx, bson.D{{"PlayerID", playerID}})
	defer cur.Close(ctx)
	if err := cur.Err(); err != nil {
		fmt.Println(err) // todo remembering print error message
		return heroList
	}

	hero := new(hero.Hero)
	for cur.Next(ctx) {
		err := cur.Decode(&hero)
		if err != nil {
			fmt.Println(err) // todo remembering print error message
			continue
		}

		heroList = append(heroList, hero)
	}

	return heroList
}

// get name of collection
// parameter：
// nothing
// result：
// 1.collection name
func getCollectionName() string {
	return "p_copy"
}
