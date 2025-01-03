package characters

import (
	"quasar/characters/actions"
	"quasar/characters/stats"

	"github.com/hajimehoshi/ebiten/v2"
)

type Character interface {
	GetStats() *stats.Statistics
	GetActionList() []*actions.Action
	Draw(*ebiten.Image, int, int)
}
