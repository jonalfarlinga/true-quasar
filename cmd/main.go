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
	// Create a new game instance
	g := game.NewGame()
	err := db.InitMySQL(getEnv())
	if err != nil {
		log.Println(err)
	}
	common.GameState = common.StateMenu

	ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
	ebiten.SetFullscreen(true)

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
