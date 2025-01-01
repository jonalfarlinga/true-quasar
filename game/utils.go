package game

import "quasar/common"

// result of rolling n action dice
// 0 - power
// 1 - accuracy
// 2 - will
func rollActionDice(n int) map[int]int {
	choices := make(map[int]int)
	choices[0] = 0
	choices[1] = 0
	choices[2] = 0
	for i := 0; i < n; i++ {
		choices[common.RandInt(1, 3)]++
	}
	return choices
}