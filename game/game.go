package game

import (
	"fmt"
	"quasar/characters"
	"quasar/common"
	"quasar/db"
	"quasar/game/combat"
	"quasar/game/draft"
	"quasar/game/menu"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Team      *characters.Team
	OpFor     *characters.OpFor
	TurnOrder []*characters.Hero
}

func NewGame() *Game {
	return &Game{
		Team:      characters.NewTeam(),
		OpFor:     &characters.OpFor{},
		TurnOrder: make([]*characters.Hero, 4),
	}
}

func (g *Game) Update() error {
	switch common.GameState {
	case common.StateMenu:
		menu.Update()
	case common.StateDraft:
		g.OpFor = characters.DefaultOpFor()
		draft.Update(g.Team)
	case common.StateCombat:
		combat.Update(g.Team, g.OpFor)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(common.BackgroundColor)

	switch common.GameState {
	case common.StateMenu:
		menu.Draw(screen)
	case common.StateDraft:
		draft.Draw(screen, g.Team)
	case common.StateCombat:
		combat.Draw(screen, g.Team, g.OpFor)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%+v", db.Pool))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth, common.ScreenHeight
}
