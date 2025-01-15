package combatdata

import "quasar/characters"

type TurnManager struct {
	ActiveChar      *characters.Character
	OffenseTargets  []*characters.Character
	FriendlyTargets []*characters.Character
}

func NewTurnManager() *TurnManager {
	return &TurnManager{
		ActiveChar:      nil,
		OffenseTargets:  make([]*characters.Character, 0),
		FriendlyTargets: make([]*characters.Character, 0),
	}
}
