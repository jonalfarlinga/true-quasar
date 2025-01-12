package combat

import (
	"time"

	cd "quasar/game/combat/combatdata"
)

func start_turn() {
	turnStartTime = time.Now()
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

func animate() {
	if done := cd.AnimationsDone(); done {
		combatState = CombatStateContinue
	}
}
