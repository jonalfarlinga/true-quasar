package main

import (
	"log"
	"os"
	"quasar/common"
	"quasar/db"
	"quasar/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize background processes
	err := db.InitMySQL(getEnv())
	if err != nil {
		log.Println(err)
	}
	common.LoadEmblems()

	// Create a new game instance
	g := game.NewGame()
	common.GameState = common.StateMenu

	// Set window size and run in fullscreen
	ebiten.SetWindowSize(int(common.ScreenWidth), int(common.ScreenHeight))
	ebiten.SetFullscreen(true)

	// Set the frame rate
	ebiten.SetTPS(common.FrameRate)

	// Run the game
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

func getEnv() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("DSN")
}
