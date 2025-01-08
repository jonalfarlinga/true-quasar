package characters

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type OpFor struct {
	Chars []Character
}

func DefaultOpFor() *OpFor {
	of := &OpFor{
		Chars: make([]Character, 1),
	}
	of.Chars[0] = DefaultBoss()
	return of
}

func (of *OpFor) DrawCombatHeader(screen *ebiten.Image) {
	char := of.Chars[0]
	boss := char.(*Enemy)
	boss.DrawCombatHeader(screen)
}
