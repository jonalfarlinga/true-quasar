package characters

import (
	"quasar/characters/actions"
	"quasar/characters/stats"

	"github.com/hajimehoshi/ebiten/v2"
)

type Character interface {
	GetStats() *stats.Statistics
	GetActionList() []*actions.Action
	DrawChar(*ebiten.Image, int, int)
	DrawIcon(*ebiten.Image, int, int)
	DrawTooltip(*ebiten.Image, int, int)
	GetType() uint8
}
