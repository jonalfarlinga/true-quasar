package characters

import (
	"quasar/characters/actions"
	"quasar/characters/stats"

	"github.com/hajimehoshi/ebiten/v2"
)

type Character interface {
	EffectiveStats() *stats.Statistics
	GetTurnmeter() int
	SetTurnmeter(int)
	Tick() int
	GetActionList() []*actions.Action
	GetName() string
	GetAvatar() *ebiten.Image
	DrawAvatar(*ebiten.Image, int, int)
	DrawIcon(*ebiten.Image, int, int)
	DrawTooltip(*ebiten.Image, int, int)
	GetType() uint8
}
