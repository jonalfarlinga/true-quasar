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

func NewCombatData() *CombatData {
	game := &CombatData{
		Team:      characters.NewTeam(),
		OpFor:     characters.DefaultOpFor(),
		TurnOrder: make([]*characters.Hero, 4),
		Area:      areas.DefaultBattlefield(),
	}

	game.Team.Heroes[0] = characters.DefaultHero("vanguard", stats.DefaultStats(), actions.ActionsDefault())
	return game
}
