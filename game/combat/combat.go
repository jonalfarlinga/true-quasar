package combat

import (
	"quasar/characters"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
)

var prevMousePressed bool = true
var exitButton common.Button = common.Button{
	X:      common.ScreenWidth - 55,
	Y:      15,
	Width:  40,
	Height: 40,
	Text:   "X",
	Active: true,
}

func Update(team *characters.Team) {
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if mousePressed && !prevMousePressed {
		x, y := ebiten.CursorPosition()
		if common.Collide(x, y, &exitButton) {
			common.GameState = common.StateMenu
			team.Heroes = make([]*characters.Hero, 4)
		}
	}
	prevMousePressed = mousePressed
}

func Draw(screen *ebiten.Image) {
	screen.Fill(common.BackgroundColor)

	// draw combat screen

	exitButton.Draw(screen)
}
