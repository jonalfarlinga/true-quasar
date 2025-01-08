package combat

import (
	"quasar/characters"
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
			cd.CD.Team.Heroes = make([]*characters.Hero, 4)
		}
	}
	prevMousePressed = mousePressed
}

func Draw(screen *ebiten.Image) {
	screen.Fill(common.BackgroundColor)

	// draw combat screen
	vector.DrawFilledRect(screen, 10, 10, common.ScreenWidth-20, 65, common.PanelColor, false)
	cd.CD.OpFor.DrawCombatHeader(screen)

	// draw combat area
	vector.DrawFilledRect(screen, 10, 85, common.ScreenWidth-20, common.ScreenHeight-90, common.PanelColor, false)
    cd.CD.Area.Draw(screen, 40, 160)

    // draw characters
    cd.CD.OpFor.Chars[0].DrawHero(screen, 80, 290)
    for i, hero := range cd.CD.Team.Heroes {
        hero.DrawHero(screen, int(common.ScreenWidth*3/4)-160-(i*125), int(common.ScreenHeight-280))
    }
    cd.CD.Team.DrawIcons(screen, float64(common.ScreenWidth)-200, 160, false)

	exitButton.Draw(screen)
}
