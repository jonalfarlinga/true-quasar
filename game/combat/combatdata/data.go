package combatdata

import (
	"quasar/areas"
	"quasar/characters"
	"quasar/characters/actions"
	"quasar/characters/stats"
)

var cd *CombatData

type CombatData struct {
	Team      *characters.Team
	OpFor     *characters.OpFor
	TurnOrder *TurnTracker
	Area      *areas.Battlefield
}

func NewCombatData() {
	game := &CombatData{
		Team:      characters.NewTeam(),
		OpFor:     characters.DefaultOpFor(),
		TurnOrder: NewTurnTracker(nil),
		Area:      areas.DefaultBattlefield(),
	}

	game.Team.Heroes[0] = characters.DefaultHero("vanguard", stats.DefaultStats(), actions.ActionsDefault())
	cd = game
}

func InitTeam() {
	cd.Team.Heroes = make([]*characters.Hero, 4)
}

func DefaultTeam() {
	cd.Team = characters.NewTeam()
}

func OpFor() *characters.OpFor {
	return cd.OpFor
}

func Area() *areas.Battlefield {
	return cd.Area
}

func Boss() *characters.Enemy {
	return cd.OpFor.Chars[0].(*characters.Enemy)
}

func Team() *characters.Team {
	return cd.Team
}

func TurnOrder() *TurnTracker {
	return cd.TurnOrder
}

func InitTurnOrder() {
	cd.TurnOrder = NewTurnTracker(cd.OpFor.Chars)
	for _, hero := range cd.Team.Heroes {
		cd.TurnOrder.AddCharacter(hero)
	}
}
