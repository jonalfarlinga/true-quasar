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
	return &Hero{
		id:          id,
		Name:        name,
		Description: description,
		Type:        heroType,
		Role:        role,
		Stats:       statistics,
		ActionList:  actionList,
		Cooldowns:   make([]int, len(actionList)),
		HeroImage:   assets.MustLoadImage(fmt.Sprintf("images/%s_hero.png", name)),
		IconImage:   assets.MustLoadImage(fmt.Sprintf("images/%s_icon.png", name)),
		BannerImage: assets.MustLoadImage(fmt.Sprintf("images/%s_banner.png", name)),
	}
}

func DefaultHero(name string, statistics *stats.Statistics, actionList []*actions.Action) *Hero {
	if statistics == nil {
		statistics = stats.DefaultStats()
	}
	if len(actionList) == 0 {
		actionList = actions.ActionsDefault()
	}

	return &Hero{
		Name:        name,
		Stats:       statistics,
		Type:        uint8(rand.Intn(6)),
		Description: "A long description about a hero to test the word wrapping function.",
		ActionList:  actionList,
		Cooldowns:   make([]int, len(actionList)),
		HeroImage:   assets.MustLoadImage("images/vanguard_hero.png"),
		IconImage:   assets.MustLoadImage("images/vanguard_icon.png"),
		BannerImage: assets.MustLoadImage("images/vanguard_banner.png"),
	}
}

func (h *Hero) GetID() int {
	return h.id
}
