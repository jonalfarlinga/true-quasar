package menu

import (
	"image/color"
	"os"
	"quasar/consts"
	"quasar/state"

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
	buttonColor := color.Black
	x, y := ebiten.CursorPosition()
	if collide(x, y, m) {
		buttonColor = color.Gray16{Y: 0x7FFF}
	}
	vector.DrawFilledRect(screen, float32(m.X), float32(m.Y), float32(m.Width), float32(m.Height), buttonColor, false)

	face := basicfont.Face7x13
	text.Draw(screen, m.Text, face, m.X+10, m.Y+m.Height/2, color.White)
}

var PlayButton MenuButton = MenuButton{
	X:      consts.ScreenWidth / 2 - 100,
	Y:      consts.ScreenHeight / 2 - 125,
	Width:  200,
	Height: 50,
	Text:   "Play",
}

var CreateButton MenuButton = MenuButton{
	X:      consts.ScreenWidth / 2 - 100,
	Y:      consts.ScreenHeight / 2 - 25,
	Width:  200,
	Height: 50,
	Text:   "Create",
}

var ExitButton MenuButton = MenuButton{
	X:      consts.ScreenWidth / 2 - 100,
	Y:      consts.ScreenHeight / 2 + 75,
	Width:  200,
	Height: 50,
	Text:   "Exit",
}

func Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if collide(x, y, &PlayButton) {
			state.GameState = state.StateDraft
		} else if collide(x, y, &ExitButton) {
			os.Exit(0)
		}
	}
}

func Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	PlayButton.Draw(screen)
	CreateButton.Draw(screen)
	ExitButton.Draw(screen)
}

func collide(x, y int, button *MenuButton) bool {
	return (x > button.X &&
		x < button.X+button.Width &&
		y > button.Y &&
		y < button.Y+button.Height)
}
