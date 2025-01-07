package draft

import (
	"math/rand"
	"quasar/characters"
	"quasar/common"
	"quasar/db"
)

const (
	DraftStateFirst uint8 = iota
	DraftStateSecond
	DraftStateThird
	DraftStateFourth
	DraftStateDone
	DraftStateStart
)

var draftLineUp []*DraftCard
var draftState uint8 = DraftStateStart
var draftSelection = -1
var saveSelection = -1

var exitButton common.Button = common.Button{
	X:      common.ScreenWidth - 55,
	Y:      15,
	Width:  40,
	Height: 40,
	Text:   "X",
	Active: true,
}
var lockInButton common.Button = common.Button{
	X:      common.ScreenWidth/2 - 100,
	Y:      common.ScreenHeight - 75,
	Width:  200,
	Height: 50,
	Text:   "Lock In",
	Active: false,
}
var fightButton common.Button = common.Button{
	X:      common.ScreenWidth/2 - 100,
	Y:      common.ScreenHeight - 75,
	Width:  200,
	Height: 50,
	Text:   "Fight",
	Active: false,
}

func getDraftPool() {
	n := float32(len(draftLineUp)) / 2
	mod := float32((len(draftLineUp) % 2) * 5)
	x := int(float32(common.ScreenWidth/2) - (180 * n) + mod)
	y := 225
	list := getHeroes()
	for i := 0; i < len(draftLineUp); i++ {
		if i < 4 {
			// a := actions.ActionsDefault()
			// s := stats.DefaultStats()
			// n := "Hero " + strconv.Itoa(rand.Intn(100))
			draftLineUp[i] = &DraftCard{
				Hero:   list[i],
				Active: true,
			}
		}
		draftLineUp[i].SetPosition(x, y)
		x += 180
	}
}

func getHeroes() []*characters.Hero {
	var heroes []*characters.Hero
	var err error
	switch draftState {
	case DraftStateStart:
		// get all defenders
		heroes, err = db.GetHeroesByRole(db.Pool, common.Defender)
	case DraftStateFirst:
		// get all strikers
		heroes, err = db.GetHeroesByRole(db.Pool, common.Striker)
	case DraftStateSecond:
		// get all controllers
		heroes, err = db.GetHeroesByRole(db.Pool, common.Controller)
	case DraftStateThird:
		// get all channelers
		heroes, err = db.GetHeroesByRole(db.Pool, common.Channeler)
	default:

	}
	if err != nil {
		return nil
	}
	rand.Shuffle(len(heroes), func(i, j int) {
		heroes[i], heroes[j] = heroes[j], heroes[i]
	})
	return heroes
}
