package combat

import (
	"quasar/common"
	cd "quasar/game/combat/combatdata"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

func Update() {
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if mousePressed && !prevMousePressed {
		x, y := ebiten.CursorPosition()
		if common.Collide(x, y, &exitButton) {
			common.GameState = common.StateMenu
			cd.InitTeam()
		}
	}
	prevMousePressed = mousePressed
}

func Draw(screen *ebiten.Image) {
	screen.Fill(common.BackgroundColor)

	// draw combat screen
	vector.DrawFilledRect(screen, 10, 10, common.ScreenWidth-20, 115, common.PanelColor, false)
	cd.OpFor().DrawCombatHeader(screen)

	// draw combat area
	vector.DrawFilledRect(screen, 10, 145, common.ScreenWidth-20, common.ScreenHeight-90, common.PanelColor, false)
	cd.Area().Draw(screen, 40, 160)

	// draw characters
	cd.Boss().DrawChar(screen, 80, 290)
	for i, hero := range cd.Team().Heroes {
		hero.DrawChar(screen, int(common.ScreenWidth*3/4)-160-(i*125), int(common.ScreenHeight-280))
	}
	cd.Team().DrawIcons(screen, float64(common.ScreenWidth)-130, 160, false)

	exitButton.Draw(screen)
}
