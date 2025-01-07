package menu

import (
	"os"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
)

var PlayButton common.Button = common.Button{
	X:      common.ScreenWidth/2 - 100,
	Y:      common.ScreenHeight/2 - 125,
	Width:  200,
	Height: 50,
	Text:   "Play",
	Active: true,
}
var CreateButton common.Button = common.Button{
	X:      common.ScreenWidth/2 - 100,
	Y:      common.ScreenHeight/2 - 25,
	Width:  200,
	Height: 50,
	Text:   "Create",
	Active: true,
}
var ExitButton common.Button = common.Button{
	X:      common.ScreenWidth/2 - 100,
	Y:      common.ScreenHeight/2 + 75,
	Width:  200,
	Height: 50,
	Text:   "Exit",
	Active: true,
}

func Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if common.Collide(x, y, &PlayButton) {
			common.GameState = common.StateDraft
		} else if common.Collide(x, y, &CreateButton) {
			common.GameState = common.StateCombat
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
