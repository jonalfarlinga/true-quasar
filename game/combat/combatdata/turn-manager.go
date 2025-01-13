package combatdata

import "quasar/characters"

type TurnManager struct {
	ActiveChar      *characters.Character
	OffenseTargets  []*characters.Character
	FriendlyTargets []*characters.Character
}
