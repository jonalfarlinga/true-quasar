package game

import (
	"quasar/characters"
	"quasar/consts"
	"quasar/state"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Team      []Team
	TurnOrder []*characters.Hero
}

type Team struct {
	Slot1       *characters.Hero
	Slot2       *characters.Hero
	Slot3       *characters.Hero
	Slot4       *characters.Hero
	DamageDealt []int
	DamageTaken []int
	HealsDealt  []int
}

func (g *Game) Update() error {
	switch state.GameState {
	case consts.StateMenu:
		// Menu logic goes here
	case consts.StateDraft:
		// Draft logic goes here
	case consts.StateCombat:
		// Combat logic goes here
	}

	// Game logic goes here
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch state.GameState {
	case consts.StateMenu:
		// Draw menu screen
	case consts.StateDraft:
		// Draw draft screen
	case consts.StateCombat:
		// Draw combat screen
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the game screen size
	return outsideWidth, outsideHeight
}
