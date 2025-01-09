package characters

import (
	"fmt"
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

type Hero struct {
	id          int
	Name        string
	Description string
	Type        uint8
	Role        uint8
	Stats       *stats.Statistics
	EffStats    *stats.Statistics
	Turnmeter   int
	ActionList  []*actions.Action
	Cooldowns   []int
	AvatarImage *ebiten.Image
	IconImage   *ebiten.Image
	BannerImage *ebiten.Image
}

func (h *Hero) GetID() int {
	return h.id
}

func (h *Hero) EffectiveStats() *stats.Statistics {
	return h.EffStats
}

func (e *Hero) GetTurnmeter() int {
	return e.Turnmeter
}

func (e *Hero) Tick() int {
	e.Turnmeter += e.Stats.Speed
	return e.Turnmeter
}

func (h *Hero) GetActionList() []*actions.Action {
	return h.ActionList
}

func (h *Hero) GetType() uint8 {
	return h.Type
}

func (h *Hero) GetAvatar() *ebiten.Image {
	return h.AvatarImage
}

func (h *Hero) DrawAvatar(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	if h.AvatarImage == nil {
		screen.DrawImage(assets.Placeholder, op)
	}
	screen.DrawImage(h.AvatarImage, op)
}

func (h *Hero) DrawIcon(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(h.IconImage, op)
	text.Draw(screen, h.Name, common.MenuFont, int(op.GeoM.Element(0, 2))+10, int(op.GeoM.Element(1, 2))+110, color.White)
	op.GeoM.Reset()
	op.GeoM.Scale(0.25, 0.25)
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(common.Emblems[h.Type], op)
	op.GeoM.Translate(75, 0)
	screen.DrawImage(common.Emblems[h.Role+6], op)
}

func (h *Hero) DrawTooltip(screen *ebiten.Image, x, y int) {
	const ttWidth, ttHeight float32 = 300, 210
	// const textWidth = 39
	X, Y := float32(x), float32(y)
	if X+ttWidth > common.ScreenWidth {
		X -= ttWidth
	}
	if Y+ttHeight > common.ScreenHeight {
		Y -= ttHeight
	}
	vector.DrawFilledRect(screen, X, Y, ttWidth, ttHeight, common.BackgroundColor, false)
	typerole := fmt.Sprintf("%s - %s", common.HeroTypeNames[h.Type], common.HeroRoleNames[h.Role])
	text.Draw(screen, h.Name, common.MenuFont, int(X+10), int(Y+18), color.White)
	text.Draw(screen, typerole, common.MenuFont, int(X+10), int(Y+33), color.White)
	text.Draw(screen, h.Stats.ToString(), common.MenuFont, int(X+10), int(Y+48), color.White)

}

func NewHero(id int, name string, description string, heroType uint8, role uint8, statistics *stats.Statistics, actionList []*actions.Action) *Hero {
	heroImg, err := assets.MustLoadImage(fmt.Sprintf("images/avatar/%s.png", name))
	if err != nil {
		heroImg, _ = assets.MustLoadImage("images/placeholder.png")
	}
	iconImg, err := assets.MustLoadImage(fmt.Sprintf("images/icon/%s.png", name))
	if err != nil {
		iconImg, _ = assets.MustLoadImage("images/placeholder.png")
	}
	bannerImg, err := assets.MustLoadImage(fmt.Sprintf("images/banner/%s.png", name))
	if err != nil {
		bannerImg, _ = assets.MustLoadImage("images/placeholder.png")
	}
	return &Hero{
		id:          id,
		Name:        name,
		Description: description,
		Type:        heroType,
		Role:        role,
		Stats:       statistics,
		EffStats:    statistics,
		ActionList:  actionList,
		Cooldowns:   make([]int, len(actionList)),
		AvatarImage: heroImg,
		IconImage:   iconImg,
		BannerImage: bannerImg,
	}
}

func DefaultHero(name string, statistics *stats.Statistics, actionList []*actions.Action) *Hero {
	if statistics == nil {
		statistics = stats.DefaultStats()
	}
	if len(actionList) == 0 {
		actionList = actions.ActionsDefault()
	}
	heroImg, err := assets.MustLoadImage("images/avatar/vanguard.png")
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

	return &Hero{
		Name:        name,
		Stats:       statistics,
		EffStats:    statistics,
		Type:        uint8(rand.Intn(6)),
		Description: "A long description about a hero to test the word wrapping function.",
		ActionList:  actionList,
		Cooldowns:   make([]int, len(actionList)),
		AvatarImage: heroImg,
		IconImage:   iconImg,
		BannerImage: bannerImg,
	}
}
