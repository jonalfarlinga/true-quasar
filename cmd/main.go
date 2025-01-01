package main

import (
	"quasar/game"
	"quasar/state"
	"quasar/consts"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Create a new game instance
	g := &game.Game{}
	state.GameState = state.StateMenu

	ebiten.SetWindowSize(consts.ScreenWidth, consts.ScreenHeight)
	ebiten.SetFullscreen(true)

	// Run the game
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
