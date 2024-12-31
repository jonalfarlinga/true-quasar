package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

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

func main() {
	// Create a new game instance
	game := &Game{}

	// Set the window to full-screen mode
	ebiten.SetFullscreen(true)

	// Run the game
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
