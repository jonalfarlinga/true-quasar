package menu

import (
	"image/color"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

type MenuButton struct {
	X, Y, Width, Height int
	Text                string
}

func (m *MenuButton) Draw(screen *ebiten.Image) {
	buttonColor := common.MenuButtonColor
	x, y := ebiten.CursorPosition()
	if common.Collide(x, y, m) {
		buttonColor = common.MenuButtonHoverColor
	}
	vector.DrawFilledRect(screen, float32(m.X), float32(m.Y), float32(m.Width), float32(m.Height), buttonColor, false)

	face := basicfont.Face7x13
	text.Draw(screen, m.Text, face, m.X+10, m.Y+m.Height/2, color.White)
}

func (m *MenuButton) GetBounds() (int, int, int, int) {
	return m.X, m.Y, m.X + m.Width, m.Y + m.Height
}
