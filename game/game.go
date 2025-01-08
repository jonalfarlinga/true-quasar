package game

import (
	"fmt"
	"quasar/common"
	"quasar/db"
	"quasar/game/combat"
	cd "quasar/game/combat/combatdata"
	"quasar/game/draft"
	"quasar/game/menu"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
}

func NewGame() *Game {
	cd.CD = cd.NewCombatData()
	return &Game{}
}

func (g *Game) Update() error {
	switch common.GameState {
	case common.StateMenu:
		menu.Update()
	case common.StateDraft:
		draft.Update()
	case common.StateCombat:
		combat.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(common.BackgroundColor)

	switch common.GameState {
	case common.StateMenu:
		menu.Draw(screen)
	case common.StateDraft:
		draft.Draw(screen)
	case common.StateCombat:
		combat.Draw(screen)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%+v", db.Pool))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(common.ScreenWidth), int(common.ScreenHeight)
}
