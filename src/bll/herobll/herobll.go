package herobll

import (
	"fmt"
	"moqikaka.com/mongodb_go/src/dal/herodal"
	"moqikaka.com/mongodb_go/src/model/hero"
	"strconv"
	"time"
)

// init data
// parameter：
// playerID：unique id
// result：nothing
func initData(playerID string) {
	heroMap := make(map[string]*hero.Hero, 10)

	heroList := herodal.GetList(playerID)
	for _, item := range heroList {
		if _, exists := heroMap[item.HeroID]; exists {
			continue
		}

		heroMap[item.HeroID] = item
	}
}

func GetData(playerID string) []*hero.Hero {
	initData(playerID)

	return herodal.GetList(playerID)
}

// update info
func UpdateInfo() {
	for j := 1; j <= 100000; j++ {
		go func() {
			for i := 1; i <= 100000; i++ {
				modelID, _ := strconv.Atoi(fmt.Sprintf("1201%d", j))
				hero := hero.NewHero(fmt.Sprintf("04183f22-d990-4d5c-9d1a-5ef4e6800088_%d", j), fmt.Sprintf("25beceb6-10d7-4310-bf88-e55080e50c6e_%d", j), modelID, j, j)
				heroList := make([]interface{}, 0, 10)
				heroList = append(heroList, hero)

				herodal.InsertData(heroList)
			}
		}()
	}

	time.Sleep(10 * time.Minute)
}

// update info
func UpdateInfo1() {
	hero := hero.NewHero1("04183f22-d990-4d5c-9d1a-5ef4e6800088tpx", "25beceb6-10d7-4310-bf88-e55080e50c6etps", 22222, 100, 100, 20000)
	heroList := make([]interface{}, 0, 10)
	heroList = append(heroList, hero)

	herodal.InsertData(heroList)
}
