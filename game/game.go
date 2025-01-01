package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{
	Team []Team
	TurnOrder []*Character
}

type Team struct {
	Slot1 *Character
	Slot2 *Character
	Slot3 *Character
	Slot4 *Character
	DamageDealt []int
	DamageTaken []int
	HealsDealt []int
}

func (g *Game) Update() error {
	// Game logic goes here
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Rendering logic goes here
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the game screen size
	return outsideWidth, outsideHeight
}
