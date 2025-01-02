package main

import (
	"quasar/common"
	"quasar/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Create a new game instance
	g := game.NewGame()
	common.GameState = common.StateMenu
	// draft.DraftState = draft.DraftStateFirst

	ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
	ebiten.SetFullscreen(true)

	// Run the game
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
