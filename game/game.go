package game

import (
	"quasar/characters"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	State	  uint8
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
	switch g.State {
	case StateMenu:
		// Menu logic goes here
	case StateDraft:
		// Draft logic goes here
	case StateCombat:
		// Combat logic goes here
	}

	// Game logic goes here
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.State {
	case StateMenu:
		// Draw menu screen
	case StateDraft:
		// Draw draft screen
	case StateCombat:
		// Draw combat screen
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the game screen size
	return outsideWidth, outsideHeight
}
