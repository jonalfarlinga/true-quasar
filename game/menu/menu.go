package menu

import (
	"image/color"
	"os"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)


type MenuButton struct {
	X, Y, Width, Height int
	Text                string
}

func (m *MenuButton) Draw(screen *ebiten.Image) {
	buttonColor := color.Gray16{Y: 0x555F}
	x, y := ebiten.CursorPosition()
	if common.Collide(x, y, m) {
		buttonColor = color.Gray16{Y: 0x7FFF}
	}
	vector.DrawFilledRect(screen, float32(m.X), float32(m.Y), float32(m.Width), float32(m.Height), buttonColor, false)

	face := basicfont.Face7x13
	text.Draw(screen, m.Text, face, m.X+10, m.Y+m.Height/2, color.White)
}

func (m *MenuButton) GetBounds() (int, int, int, int) {
	return m.X, m.Y, m.X + m.Width, m.Y + m.Height
}

var PlayButton MenuButton = MenuButton{
	X:      common.ScreenWidth / 2 - 100,
	Y:      common.ScreenHeight / 2 - 125,
	Width:  200,
	Height: 50,
	Text:   "Play",
}

var CreateButton MenuButton = MenuButton{
	X:      common.ScreenWidth / 2 - 100,
	Y:      common.ScreenHeight / 2 - 25,
	Width:  200,
	Height: 50,
	Text:   "Create",
}

var ExitButton MenuButton = MenuButton{
	X:      common.ScreenWidth / 2 - 100,
	Y:      common.ScreenHeight / 2 + 75,
	Width:  200,
	Height: 50,
	Text:   "Exit",
}

func Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if common.Collide(x, y, &PlayButton) {
			common.GameState = common.StateDraft
		} else if common.Collide(x, y, &ExitButton) {
			os.Exit(0)
		}
	}
}

func Draw(screen *ebiten.Image) {
	screen.Fill(common.BackgroundColor)
	PlayButton.Draw(screen)
	CreateButton.Draw(screen)
	ExitButton.Draw(screen)
}
