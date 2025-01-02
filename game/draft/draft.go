package draft

import (
	"image/color"
	"quasar/assets"
	"quasar/characters"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var panelColor = color.RGBA{0x30, 0x30, 0xA0, 0xFF}
var placeholder = assets.MustLoadImage("images/placeholder.png")

func Update(team *characters.Team) {
	if draftState == DraftStateStart {
		draftSetup()
		draftState = DraftStateFirst
		return
	}
	// Draft logic goes here
}

func Draw(screen *ebiten.Image, team *characters.Team) {
	screen.Fill(common.BackgroundColor)
	vector.DrawFilledRect(screen, 10, 10, common.ScreenWidth-20, 180, panelColor, false)
	vector.DrawFilledRect(screen, 10, 210, common.ScreenWidth-20, common.ScreenHeight - 220, panelColor, false)

	// Draw team
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(common.ScreenWidth)/2 - 315, 25)
	for i := 0; i < 4; i++ {
		hero := team.Heroes[i]
		if hero != nil {
			screen.DrawImage(hero.IconImage, op)
		} else {
			screen.DrawImage(placeholder, op)
		}
		op.GeoM.Translate(160, 0)
	}
	// Draw hero pool
	for i, card := range draftLineUp {
		card.Draw(screen, i == draftSelection, i == saveSelection)
	}

}
