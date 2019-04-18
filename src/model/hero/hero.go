package hero

// hero object
type Hero struct {
	// 玩家ID
	PlayerID string

	// 武将模型ID
	HeroID string

	// 模型Id
	MoldeID int

	// 等级
	Lv int

	// 阶段
	Step int

	// total experience
	ExpTotal int
}

// create new hero object
// parameter：
// playerID：unique id
// heroID：unique id of hero
// modelID： unique id of hero model
// lv：lv of hero
// step：step of hero
func NewHero(playerID, heroID string, modelID, lv, step int) *Hero {
	return &Hero{
		PlayerID: playerID,
		HeroID:   heroID,
		MoldeID:  modelID,
		Lv:       lv,
		Step:     step,
	}
}

// create new hero object
// parameter：
// playerID：unique id
// heroID：unique id of hero
// modelID： unique id of hero model
// lv：lv of hero
// step：step of hero
// experience：hero current experience
func NewHero1(playerID, heroID string, modelID, lv, step, experience int) *Hero {
	return &Hero{
		PlayerID: playerID,
		HeroID:   heroID,
		MoldeID:  modelID,
		Lv:       lv,
		Step:     step,
		ExpTotal: experience,
	}
}
