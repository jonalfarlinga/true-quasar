package characters

import (
	"fmt"
	"math/rand"
	"quasar/assets"
	"quasar/characters/actions"
	"quasar/characters/stats"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hero struct {
	id          int
	Name        string
	Description string
	Type        uint8
	Role        uint8
	Stats       *stats.Statistics
	ActionList  []*actions.Action
	Cooldowns   []int
	HeroImage   *ebiten.Image
	IconImage   *ebiten.Image
	BannerImage *ebiten.Image
}

func (h *Hero) GetStats() *stats.Statistics {
	return h.Stats
}

func (h *Hero) GetActionList() []*actions.Action {
	return h.ActionList
}

func (h *Hero) Draw(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(h.HeroImage, op)
}

func NewHero(id int, name string, description string, heroType uint8, role uint8, statistics *stats.Statistics, actionList []*actions.Action) *Hero {
	heroImg, err := assets.MustLoadImage(fmt.Sprintf("images/hero/%s.png", name))
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
		ActionList:  actionList,
		Cooldowns:   make([]int, len(actionList)),
		HeroImage:   heroImg,
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

	return &Hero{
		Name:        name,
		Stats:       statistics,
		Type:        uint8(rand.Intn(6)),
		Description: "A long description about a hero to test the word wrapping function.",
		ActionList:  actionList,
		Cooldowns:   make([]int, len(actionList)),
		HeroImage:   heroImg,
		IconImage:   iconImg,
		BannerImage: bannerImg,
	}
}

func (h *Hero) GetID() int {
	return h.id
}
