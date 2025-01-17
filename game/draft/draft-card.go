package draft

import (
	"image/color"
	"quasar/characters"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type DraftCard struct {
	Hero   *characters.Hero
	x, y   float32
	Active bool
}

func (dc *DraftCard) Draw(screen *ebiten.Image, drafted, saved bool) {
	if drafted {
		vector.DrawFilledRect(screen, float32(dc.x), float32(dc.y), 170, 450, common.GoldColor, false)
	} else if saved {
		vector.DrawFilledRect(screen, float32(dc.x), float32(dc.y), 170, 450, common.RedColor, false)
	} else {
		vector.DrawFilledRect(screen, float32(dc.x), float32(dc.y), 170, 450, common.BackgroundColor, false)
	}

	op := &ebiten.DrawImageOptions{}
	scale := 150 / float64(dc.Hero.BannerImage.Bounds().Dx())
	bannerHeight := int(float64(dc.Hero.BannerImage.Bounds().Dy()) * scale)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(dc.x+10), float64(dc.y+35))
	screen.DrawImage(dc.Hero.BannerImage, op)

	text.Draw(screen, dc.Hero.Name, common.MenuFont, int(op.GeoM.Element(0, 2))+35, int(op.GeoM.Element(1, 2))-2, color.White)
	text.Draw(screen, common.HeroTypeNames[dc.Hero.Type], common.MenuFont, int(op.GeoM.Element(0, 2))+10, int(op.GeoM.Element(1, 2))+bannerHeight+22, color.White)
	text.Draw(screen, common.HeroRoleNames[dc.Hero.Role], common.MenuFont, int(op.GeoM.Element(0, 2))+10, int(op.GeoM.Element(1, 2))+bannerHeight+34, color.White)
	text.Draw(screen, common.WrapText(dc.Hero.Description, 20), common.MenuFont, int(op.GeoM.Element(0, 2))+10, int(op.GeoM.Element(1, 2))+bannerHeight+46, color.White)

	op.GeoM.Reset()
	op.GeoM.Scale(0.25, 0.25)
	op.GeoM.Translate(float64(dc.x+10), float64(dc.y+10))
	screen.DrawImage(common.Emblems[dc.Hero.Type], op)
	op.GeoM.Translate(125, 0)
	screen.DrawImage(common.Emblems[dc.Hero.Role+6], op)
}

func (dc *DraftCard) GetBounds() (float32, float32, float32, float32) {
	// Return the bounds of the card
	return dc.x, dc.y, dc.x + 170, dc.y + 450
}

func (dc *DraftCard) SetPosition(x, y float32) {
	dc.x, dc.y = x, y
}

func (dc *DraftCard) SetActive(active bool) {
	dc.Active = active
}
