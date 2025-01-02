package characters

import (
	"quasar/characters/actions"
	"quasar/characters/stats"
)

type Character interface {
	GetStats() *stats.Statistics
	GetActionList() []*actions.Action
}
