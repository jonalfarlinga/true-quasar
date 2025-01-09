package combat

import (
	"quasar/characters"
	"quasar/common"
	cd "quasar/game/combat/combatdata"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var prevMousePressed bool = true

func Update() {
	if combatState == CombatStateStart {
		iconPanel = *characters.NewIconPanel(common.ScreenWidth-130, 160, false, 0)
		for _, hero := range cd.Team().Heroes {
			iconPanel.AddIcon(hero)
		}
		combatState = CombatStateContinue
		bossIcon.Char = cd.Boss()
	}
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if mousePressed && !prevMousePressed {
		x, y := ebiten.CursorPosition()
		if common.Collide(x, y, &exitButton) {
			common.GameState = common.StateMenu
			combatState = CombatStateStart
			cd.NewCombatData()
		}
	}
	prevMousePressed = mousePressed
}

func Draw(screen *ebiten.Image) {
	// background layer
	screen.Fill(common.BackgroundColor)
	drawPanel(screen)

	// info layer
	cd.OpFor().DrawCombatHeader(screen)
	cd.Area().DrawArea(screen, 40, 160)
	cd.Boss().DrawChar(screen, 80, 290)
	for i, hero := range cd.Team().Heroes {
		hero.DrawChar(screen, int(common.ScreenWidth*3/4)-160-(i*125), int(common.ScreenHeight-280))
	}
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
