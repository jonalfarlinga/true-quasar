package common

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

type Button struct {
	Width, Height int
	X, Y          int
	Text          string
	Active        bool
}

func (m *Button) Draw(screen *ebiten.Image) {
	buttonColor := ButtonColor
	x, y := ebiten.CursorPosition()
	if !m.Active {
		buttonColor = ButtonOffColor
	} else if Collide(x, y, m) {
		buttonColor = ButtonHoverColor
	}
	vector.DrawFilledRect(screen, float32(m.X), float32(m.Y), float32(m.Width), float32(m.Height), buttonColor, false)

	bounds := font.MeasureString(MenuFont, m.Text)
	textX := m.X + (m.Width-bounds.Floor())/2
	text.Draw(screen, m.Text, MenuFont, textX, m.Y+m.Height/2+5, color.White)
}

func (m *Button) GetBounds() (int, int, int, int) {
	return m.X, m.Y, m.X + m.Width, m.Y + m.Height
}

func (m *Button) SetActive(active bool) {
	m.Active = active
}
