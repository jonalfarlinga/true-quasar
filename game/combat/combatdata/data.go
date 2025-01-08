package combatdata

import (
	"quasar/areas"
	"quasar/characters"
	"quasar/characters/actions"
	"quasar/characters/stats"
)

var CD *CombatData

type CombatData struct {
	Team      *characters.Team
	OpFor     *characters.OpFor
	TurnOrder []*characters.Hero
	Area      *areas.Battlefield
}

func NewCombatData() {
	game := &CombatData{
		Team:      characters.NewTeam(),
		OpFor:     characters.DefaultOpFor(),
		TurnOrder: make([]*characters.Hero, 4),
		Area:      areas.DefaultBattlefield(),
	}

	game.Team.Heroes[0] = characters.DefaultHero("vanguard", stats.DefaultStats(), actions.ActionsDefault())
	CD = game
}

func InitTeam() {
	CD.Team.Heroes = make([]*characters.Hero, 4)
}

func OpFor() *characters.OpFor {
	return CD.OpFor
}

func Area() *areas.Battlefield {
	return CD.Area
}

func Boss() *characters.Enemy {
	return CD.OpFor.Chars[0].(*characters.Enemy)
}

func Team() *characters.Team {
	return CD.Team
}
