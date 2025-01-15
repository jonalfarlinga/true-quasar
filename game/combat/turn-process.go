package combat

import (
	cd "quasar/game/combat/combatdata"

	"github.com/hajimehoshi/ebiten/v2"
)

func start_turn() {
	manager = cd.NewTurnManager()
}

func action_select() {

}

func target_select() {

}

func pre_action() {

}

func process_action() {

}

func resolve_action() {

}

func end_turn() {
	cd.TurnOrder().TurnComplete(0)
}

func cleanup() {
	if done := cd.AnimationsDone(); done {
		combatState = CombatStateContinue
	}
}

func animate(screen *ebiten.Image) {
	if done := cd.AnimationsDone(); !done {
		cd.Animate(screen)
	}
}
