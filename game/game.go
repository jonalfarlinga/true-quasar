package game

import (
	"quasar/characters"
	"quasar/common"
	"quasar/game/draft"
	"quasar/game/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Team      *characters.Team
	TurnOrder []*characters.Hero
}

func NewGame() *Game {
	return &Game{
		Team:      characters.NewTeam(),
		TurnOrder: make([]*characters.Hero, 4),
	}
}

func (g *Game) Update() error {
	switch common.GameState {
	case common.StateMenu:
		menu.Update()
	case common.StateDraft:
		draft.Update(g.Team)
	case common.StateCombat:
		// Combat logic goes here
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
		// Draw combat screen
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth, common.ScreenHeight
}
