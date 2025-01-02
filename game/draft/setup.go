package draft

import (
	"quasar/characters"
	"quasar/characters/actions"
	"quasar/characters/stats"
	"quasar/common"
	"strconv"
)

const (
	DraftStateStart uint8 = iota
	DraftStateFirst
	DraftStateSecond
	DraftStateThird
	DraftStateFourth
)

var draftLineUp []*DraftCard
var draftState uint8 = DraftStateStart
var draftSelection = -1
var saveSelection = -1

func draftSetup() {
	draftLineUp = make([]*DraftCard, 4)
	x := common.ScreenWidth/2 - 355
	y := 250
	for i := 0; i < 4; i++ {
		a := actions.ActionsDefault()
		s := stats.StatsDefault()
		n := "Hero " + strconv.Itoa(i)
		draftLineUp[i] = &DraftCard{
			Hero: characters.NewHero(n, s, a),
		}
		draftLineUp[i].SetPosition(x, y)
		x += 180
	}
}
