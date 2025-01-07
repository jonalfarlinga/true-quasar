package characters

import (
	"image/color"
	"math/rand"
	"quasar/assets"
	"quasar/characters/actions"
	"quasar/characters/stats"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Enemy struct {
	id          int
	Name        string
	Description string
	Type        uint8
	Stats       *stats.Statistics
	ActionList  []*actions.Action
	Cooldowns   []int
	HeroImage   *ebiten.Image
	IconImage   *ebiten.Image
	BannerImage *ebiten.Image
}

func (e *Enemy) GetID() int {
	return e.id
}

func (e *Enemy) GetStats() *stats.Statistics {
	return e.Stats
}

func (e *Enemy) GetActionList() []*actions.Action {
	return e.ActionList
}

func (e *Enemy) GetType() uint8 {
	return e.Type
}

func (e *Enemy) DrawHero(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(e.HeroImage, op)
}

func (e *Enemy) DrawCombatHeader(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.25, 0.25)
	op.GeoM.Translate(common.ScreenWidth-53, 15)
	screen.DrawImage(common.Emblems[e.Type], op)
	vector.DrawFilledRect(screen, 15, 15, 1194, 25, common.ButtonOffColor, false)
	resPercent := 100 * float32(e.Stats.ResCurrent) / float32(e.Stats.Resilience)
	vector.DrawFilledRect(screen, 18, 18, 1184*resPercent, 19, common.RedColor, false)
	op.GeoM.Reset()
	op.GeoM.Translate(20, 45)
	text.Draw(screen, e.Name, common.MenuFont, int(op.GeoM.Element(0, 2)), int(op.GeoM.Element(1, 2)), color.White)
}

func DefaultEnemy() *Enemy {
	heroImg, err := assets.MustLoadImage("images/hero/vanguard.png")
	if err != nil {
		heroImg, _ = assets.MustLoadImage("images/placeholder.png")
	}
	iconImg, err := assets.MustLoadImage("images/icon/vanguard.png")
	if err != nil {
		iconImg, _ = assets.MustLoadImage("images/placeholder.png")
	}
	bannerImg, err := assets.MustLoadImage("images/banner/vanguard.png")
	if err != nil {
		bannerImg, _ = assets.MustLoadImage("images/placeholder.png")
	}
	return &Enemy{
		Name:        "Bad Guy",
		Stats:       stats.DefaultStats(),
		Type:        uint8(rand.Intn(6)),
		Description: "This GUY is BAD!",
		ActionList:  actions.ActionsDefault(),
		Cooldowns:   make([]int, 1),
		HeroImage:   heroImg,
		IconImage:   iconImg,
		BannerImage: bannerImg,
	}
}
