package game

import (
	"quasar/characters"
	"quasar/characters/actions"
	"quasar/characters/stats"
	"quasar/common"
	"quasar/game/combat"
	"quasar/game/draft"
	"quasar/game/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Team      *characters.Team
	OpFor     *characters.OpFor
	TurnOrder []*characters.Hero
}

func NewGame() *Game {
	game := &Game{
		Team:      characters.NewTeam(),
		OpFor:     characters.DefaultOpFor(),
		TurnOrder: make([]*characters.Hero, 4),
	}
	game.Team.Heroes[0] = characters.DefaultHero("vanguard", stats.DefaultStats(), actions.ActionsDefault())
	return game
}

func (g *Game) Update() error {
	switch common.GameState {
	case common.StateMenu:
		menu.Update()
	case common.StateDraft:
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
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("%+v", db.Pool))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth, common.ScreenHeight
}
