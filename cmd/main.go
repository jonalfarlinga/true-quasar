package main

import (
	"quasar/game"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Create a new game instance
	g := &game.Game{}
	common.GameState = common.StateMenu

	ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
	ebiten.SetFullscreen(true)

	// Run the game
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
