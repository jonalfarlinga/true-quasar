package draft

import (
	"math/rand"
	"quasar/characters"
	"quasar/characters/actions"
	"quasar/characters/stats"
	"quasar/common"
	"strconv"
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
	X:      common.ScreenWidth - 45,
	Y:      15,
	Width:  30,
	Height: 30,
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
	for i := 0; i < len(draftLineUp); i++ {
		if i < 4 {
			a := actions.ActionsDefault()
			s := stats.StatsDefault()
			n := "Hero " + strconv.Itoa(rand.Intn(100))
			draftLineUp[i] = &DraftCard{
				Hero:   characters.NewHero(n, s, a),
				Active: true,
			}
		}
		draftLineUp[i].SetPosition(x, y)
		x += 180
	}
}
