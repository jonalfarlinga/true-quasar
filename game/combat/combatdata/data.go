package combatdata

import (
	"quasar/areas"
	"quasar/assets/animation"
	"quasar/characters"
	"quasar/characters/actions"
	"quasar/characters/stats"

	"github.com/hajimehoshi/ebiten/v2"
)

var cd *CombatData

type CombatData struct {
	Team           *characters.Team
	OpFor          *characters.OpFor
	Area           *areas.Battlefield
	TurnOrder      *TurnTracker
	AnimationQueue *animation.AnimationQueue
}

func NewCombatData() {
	game := &CombatData{
		Team:           characters.NewTeam(),
		OpFor:          characters.DefaultOpFor(),
		Area:           areas.DefaultBattlefield(),
		TurnOrder:      NewTurnTracker(nil),
		AnimationQueue: animation.NewAnimationQueue(),
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

func Tick() characters.Character {
	activateChar := cd.TurnOrder.Tick()
	return activateChar
}

func Animate(screen *ebiten.Image) {
	cd.AnimationQueue.Animate(screen)
}

func AnimationsDone() bool {
	return cd.AnimationQueue.Length() == 0
}
