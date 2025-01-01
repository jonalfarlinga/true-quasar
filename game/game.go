package game

import (
	"quasar/characters"
	"quasar/common"
	"quasar/game/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Team      []Team
	TurnOrder []*characters.Hero
}

func (g *Game) Update() error {
	switch common.GameState {
	case common.StateMenu:
		menu.Update()
	case common.StateDraft:
		// Draft logic goes here
	case common.StateCombat:
		// Combat logic goes here
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch common.GameState {
	case common.StateMenu:
		menu.Draw(screen)
	case common.StateDraft:
		// Draw draft screen
	case common.StateCombat:
		// Draw combat screen
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
