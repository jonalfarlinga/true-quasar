package combat

import (
	"quasar/common"
	cd "quasar/game/combat/combatdata"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var prevMousePressed bool = true

func Update() {
	if combatState == CombatStateSetup {
		initialize()
        combatState = CombatStateContinue
	}

    

	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if mousePressed && !prevMousePressed {
		x, y := ebiten.CursorPosition()
		if common.Collide(x, y, &exitButton) {
			common.GameState = common.StateMenu
			combatState = CombatStateSetup
			cd.NewCombatData()
		}
	}
	prevMousePressed = mousePressed
}

func Draw(screen *ebiten.Image) {
	if combatState == CombatStateSetup {
		return
	}
	// background layer
	screen.Fill(common.BackgroundColor)
	drawPanel(screen)

	// info layer
	cd.OpFor().DrawCombatHeader(screen)
	cd.Area().DrawArea(screen, 40, 220)
    avatars.Draw(screen, -1)
    bossAvatar.Draw(screen, true, false)
	iconPanel.DrawIcons(screen)

	// button layer
	exitButton.Draw(screen)

	// popup layer
	x, y := ebiten.CursorPosition()
	iconPanel.DrawTooltips(screen, x, y)
	if common.Collide(x, y, bossIcon) {
		cd.Boss().DrawTooltip(screen, x, y)
	}
}

func drawPanel(screen *ebiten.Image) {
	// header panel
	vector.DrawFilledRect(screen, 10, 10, common.ScreenWidth-20, 115, common.PanelColor, false)
	// combat panel
	vector.DrawFilledRect(screen, 10, 145, common.ScreenWidth-20, common.ScreenHeight-90, common.PanelColor, false)

}
