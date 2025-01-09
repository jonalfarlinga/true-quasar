package areas

import (
	"quasar/assets"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
)

type Battlefield struct {
	Image *ebiten.Image
}

func DefaultBattlefield() *Battlefield {
	b := &Battlefield{}
	var err error
	b.Image, err = assets.MustLoadImage("images/area/nexus_battlefield.png")
	if err != nil {
		b.Image, _ = assets.MustLoadImage("images/placeholder.png")
	}
	return b
}

func (b *Battlefield) DrawArea(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	var width, height float64 = float64(b.Image.Bounds().Dx()), float64(b.Image.Bounds().Dy())
	var targetWidth, targetHeight float64 = float64(common.ScreenWidth) * 3 / 4, 400.0
	op.GeoM.Scale(targetWidth/width, targetHeight/height)
	op.GeoM.Translate(x, y)
	screen.DrawImage(b.Image, op)
}
