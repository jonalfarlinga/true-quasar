package state

const (
	StateMenu = iota
	StateDraft
	StateCombat
)

var GameState uint8 = StateMenu
