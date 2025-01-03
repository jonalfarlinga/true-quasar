package draft

import (
	"image/color"
	"quasar/characters"
	"quasar/common"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type DraftCard struct {
	Hero   *characters.Hero
	x, y   int
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
	op.GeoM.Translate(float64(dc.x+10), float64(dc.y+15))
	screen.DrawImage(dc.Hero.BannerImage, op)
	text.Draw(screen, dc.Hero.Name, common.MenuFont, int(op.GeoM.Element(0, 2))+10, int(op.GeoM.Element(1, 2))-2, color.White)
	text.Draw(screen, common.HeroTypeNames[dc.Hero.Type], common.MenuFont, int(op.GeoM.Element(0, 2))+10, int(op.GeoM.Element(1, 2))+bannerHeight+22, color.White)
	text.Draw(screen, common.HeroRoleNames[dc.Hero.Role], common.MenuFont, int(op.GeoM.Element(0, 2))+10, int(op.GeoM.Element(1, 2))+bannerHeight+34, color.White)
	text.Draw(screen, wrapText(dc.Hero.Description, 20), common.MenuFont, int(op.GeoM.Element(0, 2))+10, int(op.GeoM.Element(1, 2))+bannerHeight+46, color.White)
}

func (dc *DraftCard) GetBounds() (int, int, int, int) {
	// Return the bounds of the card
	return dc.x, dc.y, dc.x + 170, dc.y + 450
}

func (dc *DraftCard) SetPosition(x, y int) {
	dc.x, dc.y = x, y
}

func (dc *DraftCard) SetActive(active bool) {
	dc.Active = active
}

func wrapText(text string, lineLen int) string {
	var result string
	var line string
	words := strings.Fields(text)
	for _, word := range words {
		if len(line)+len(word)+1 > lineLen {
			result += line + "\n"
			line = word
		} else {
			if line != "" {
				line += " "
			}
			line += word
		}
	}
	if line != "" {
		result += line
	}
	return result
}
