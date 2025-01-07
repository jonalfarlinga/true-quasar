package characters

import (
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type OpFor struct {
	Chars []Character
}

func DefaultOpFor() *OpFor {
	of := &OpFor{
		Chars: make([]Character, 1),
	}
	of.Chars[0] = DefaultEnemy()
	return of
}

func (of *OpFor) DrawCombatHeader(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 10, 10, common.ScreenWidth-20, 65, common.PanelColor, false)
	char := of.Chars[0]
	boss := char.(*Enemy)
	boss.DrawCombatHeader(screen)
}
