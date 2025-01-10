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
	Turnmeter   int
	AvatarImage *ebiten.Image
	IconImage   *ebiten.Image
	BannerImage *ebiten.Image
}

func (e *Enemy) GetID() int {
	return e.id
}

func (e *Enemy) GetName() string {
	return e.Name
}

func (e *Enemy) EffectiveStats() *stats.Statistics {
	return e.Stats
}

func (e *Enemy) GetTurnmeter() int {
	return e.Turnmeter
}

func (e *Enemy) SetTurnmeter(tm int) {
	e.Turnmeter = tm
}

func (e *Enemy) Tick() int {
	e.Turnmeter += e.Stats.Speed
	return e.Turnmeter
}

func (e *Enemy) GetActionList() []*actions.Action {
	return e.ActionList
}

func (e *Enemy) GetType() uint8 {
	return e.Type
}

func (e *Enemy) GetAvatar() *ebiten.Image {
	return e.AvatarImage
}

func (e *Enemy) DrawAvatar(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(e.AvatarImage, op)
}

func (e *Enemy) DrawCombatHeader(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.25, 0.25)
	op.GeoM.Translate(float64(common.ScreenWidth)-88, 22)
	screen.DrawImage(common.Emblems[e.Type], op)
	op.GeoM.Reset()
	op.GeoM.Translate(15, 15)
	screen.DrawImage(e.IconImage, op)
	vector.DrawFilledRect(screen, 120, 18, common.ScreenWidth-220, 25, common.ButtonOffColor, false)
	resPercent := float32(e.Stats.ResCurrent) / float32(e.Stats.Resilience)
	vector.DrawFilledRect(screen, 123, 21, (common.ScreenWidth-226)*resPercent, 19, common.RedColor, false)
	op.GeoM.Reset()
	op.GeoM.Translate(130, 55)
	text.Draw(screen, e.Name, common.MenuFont, int(op.GeoM.Element(0, 2)), int(op.GeoM.Element(1, 2)), color.White)
}

func (e *Enemy) DrawIcon(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(e.IconImage, op)
}

func (e *Enemy) DrawTooltip(screen *ebiten.Image, x, y int) {
	const ttWidth, ttHeight float32 = 300, 210
	const textWidth = 39
	X, Y := float32(x), float32(y)
	if X+ttWidth > common.ScreenWidth {
		X -= ttWidth
	}
	if Y+ttHeight > common.ScreenHeight {
		Y -= ttHeight
	}
	vector.DrawFilledRect(screen, X, Y, ttWidth, ttHeight, common.BackgroundColor, false)
	text.Draw(screen, e.Name, common.MenuFont, int(X+10), int(Y+18), color.White)
	text.Draw(screen, common.HeroTypeNames[e.Type], common.MenuFont, int(X+10), int(Y+33), color.White)
	text.Draw(screen, common.WrapText(e.Description, textWidth), common.MenuFont, int(X+10), int(Y+48), color.White)
}

func DefaultBoss() *Enemy {
	heroImg, err := assets.MustLoadImage("images/boss/mog_tron_hero.png")
	if err != nil {
		heroImg, _ = assets.MustLoadImage("images/placeholder.png")
	}
	iconImg, err := assets.MustLoadImage("images/boss/mog_tron_icon.png")
	if err != nil {
		iconImg, _ = assets.MustLoadImage("images/placeholder.png")
	}
	bannerImg, err := assets.MustLoadImage("images/boss/mog_tron_banner.png")
	if err != nil {
		bannerImg, _ = assets.MustLoadImage("images/placeholder.png")
	}
	return &Enemy{
		Name:        "Bad Guy",
		Stats:       stats.DefaultStats(),
		Type:        uint8(rand.Intn(6)),
		Description: "This GUY is BAD! He is so bad he'll mess you up. Take him out quick before the baddies win!",
		ActionList:  actions.ActionsDefault(),
		Cooldowns:   make([]int, 1),
		AvatarImage: heroImg,
		IconImage:   iconImg,
		BannerImage: bannerImg,
	}
}
