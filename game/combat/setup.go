package combat

import "quasar/characters"

const (
	CombatStateStart uint8 = iota
	CombatStateContinue
)

var combatState uint8 = CombatStateStart
var iconPanel characters.IconPanel
