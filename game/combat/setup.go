package combat

import (
	"quasar/characters"
	"quasar/common"
	cd "quasar/game/combat/combatdata"
)

const (
	CombatStateContinue uint8 = iota
	CombatStateStartTurn
	CombatStateActionSelect
	CombatStateTargetSelect
	CombatStatePreAction
	CombatStateProcessAction
	CombatStateResolveAction
	CombatStateEndTurn
	CombatStateAnimate
	CombatStateSetup
)

var combatState uint8 = CombatStateSetup
var iconPanel *characters.IconPanel
var avatars *characters.AvatarSquad
var bossAvatar *characters.Avatar

var exitButton common.Button = common.Button{
	X:      common.ScreenWidth - 55,
	Y:      15,
	Width:  40,
	Height: 40,
	Text:   "X",
	Active: true,
}
var bossIcon *characters.Icon

func initialize() {
	iconPanel = characters.NewIconPanel(common.ScreenWidth-130, 160, false, 0)
	avatars = characters.NewAvatarSquad(common.ScreenWidth*3/4-600, common.ScreenHeight-280)
	for _, hero := range cd.Team().Heroes {
		iconPanel.AddIcon(hero)
		avatars.AddAvatar(hero)
	}
	bossIcon = characters.NewIcon(15, 15, cd.Boss())
	bossAvatar = characters.NewAvatar(cd.Boss(), 80, 290)
	cd.InitTurnOrder()
}
