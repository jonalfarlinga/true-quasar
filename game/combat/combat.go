package combat

import (
	"fmt"
	"quasar/characters"
	"quasar/common"
	cd "quasar/game/combat/combatdata"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var prevMousePressed bool = true
var activeChar characters.Character
var turnStartTime time.Time

func Update() {
	// allow exit button to interrupt
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	x, y := ebiten.CursorPosition()
	if mousePressed && !prevMousePressed {
		if common.Collide(x, y, &exitButton) {
			common.GameState = common.StateMenu
			combatState = CombatStateSetup
			cd.NewCombatData()
		}
	}

	// update based on game progress
	switch combatState {
	case CombatStateSetup:
		initialize()
		combatState = CombatStateContinue
	case CombatStateContinue:
		activeChar = cd.Tick()
		if activeChar != nil {
			combatState++
		}
	case CombatStateStartTurn:
		start_turn()
		combatState++
	case CombatStateActionSelect:
		action_select()
		combatState++
	case CombatStateTargetSelect:
		target_select()
		combatState++
	case CombatStatePreAction:
		pre_action()
		combatState++
	case CombatStateProcessAction:
		process_action()
		combatState++
	case CombatStateResolveAction:
		resolve_action()
		combatState++
	case CombatStateEndTurn:
		end_turn()
		activeChar = nil
        combatState = CombatStateContinue
	}

	prevMousePressed = mousePressed
}

func Draw(screen *ebiten.Image) {
	if combatState == CombatStateSetup {
		return
	}
	// background layer
	screen.Fill(common.BackgroundColor)
	draw_panel(screen)

	// character layer
	cd.Area().DrawArea(screen, 40, 220)
	cd.OpFor().DrawCombatHeader(screen)
	draw_boss(screen)

	// action layer
	if combatState == CombatStateAnimate {
		// animate the actions
		animate()
	}

	// button layer
	exitButton.Draw(screen)

	// popup layer
	x, y := ebiten.CursorPosition()
	iconPanel.DrawTooltips(screen, x, y)
	if common.Collide(x, y, bossIcon) {
		cd.Boss().DrawTooltip(screen, x, y)
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(" Combat State: %d", combatState), 0, 15)
	ebitenutil.DebugPrintAt(screen, cd.TurnOrder().ToString(), 10, 30)
}

func draw_panel(screen *ebiten.Image) {
	// header panel
	vector.DrawFilledRect(screen, 10, 10, common.ScreenWidth-20, 115, common.PanelColor, false)
	// combat panel
	vector.DrawFilledRect(screen, 10, 145, common.ScreenWidth-20, common.ScreenHeight-90, common.PanelColor, false)
}

func draw_boss(screen *ebiten.Image) {
	avatars.Draw(screen, -1)
	bossAvatar.Draw(screen, activeChar == cd.OpFor().Chars[0], false)
	iconPanel.DrawIcons(screen)
}
