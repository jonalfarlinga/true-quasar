package combat

import (
	"quasar/characters"
	"quasar/common"
)

const (
	CombatStateStart uint8 = iota
	CombatStateContinue
)

var combatState uint8 = CombatStateStart
var iconPanel characters.IconPanel

var exitButton common.Button = common.Button{
	X:      common.ScreenWidth - 55,
	Y:      15,
	Width:  40,
	Height: 40,
	Text:   "X",
	Active: true,
}
var bossIcon *characters.Icon = &characters.Icon{
	X:   15,
	Y:   15,
	Char: nil,
}
