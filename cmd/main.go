package main

import (
	"quasar/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Create a new game instance
	g := &game.Game{}

	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetFullscreen(true)

	// Run the game
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
