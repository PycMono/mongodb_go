package herobll

import (
	"moqikaka.com/mongodb_go/src/dal/copydal"
	"moqikaka.com/mongodb_go/src/model/hero"
)

// init data
// parameter：
// playerID：unique id
// result：nothing
func initData(playerID string) {
	heroMap := make(map[string]*hero.Hero, 10)

	heroList := copydal.GetList(playerID)
	for _, item := range heroList {
		if _, exists := heroMap[item.HeroID]; exists {
			continue
		}

		heroMap[item.HeroID] = item
	}
}

func GetData(playerID string) []*hero.Hero {
	initData(playerID)

	return copydal.GetList(playerID)
}
