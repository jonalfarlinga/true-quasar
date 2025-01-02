package draft

import (
	"quasar/characters"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type DraftCard struct {
	Hero *characters.Hero
	x, y int
}

func (dc *DraftCard) Draw(screen *ebiten.Image, drafted, saved bool) {
	if drafted {
		vector.DrawFilledRect(screen, float32(dc.x), float32(dc.y), 170, 450, common.GoldColor, false)
	} else if saved {
		vector.DrawFilledRect(screen, float32(dc.x), float32(dc.y), 170, 450, common.RedColor, false)
	} else {
		vector.DrawFilledRect(screen, float32(dc.x), float32(dc.y), 170, 450, common.BackgroundColor, false)
	}
	// Draw the card
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(150/float64(dc.Hero.BannerImage.Bounds().Dx()), 150/float64(dc.Hero.BannerImage.Bounds().Dx()))
	op.GeoM.Translate(float64(dc.x+10), float64(dc.y+10))
	screen.DrawImage(dc.Hero.BannerImage, op)
}

func (dc *DraftCard) GetBounds() (int, int, int, int) {
	// Return the bounds of the card
	return dc.x, dc.y, dc.x + 170, dc.y + 450
}

func (dc *DraftCard) SetPosition(x, y int) {
	dc.x, dc.y = x, y
}