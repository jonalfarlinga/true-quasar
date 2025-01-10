package combat

import "time"

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

}

func animate() {
	if time.Since(turnStartTime) > 2*time.Second {
		combatState++
	}
}
