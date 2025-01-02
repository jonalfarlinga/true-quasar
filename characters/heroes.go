package characters

import (
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

func NewHero(id int, name string, description string, heroType uint8, role uint8, statistics *stats.Statistics, actionList []*actions.Action, heroImage, iconImage, bannerImage string) *Hero {
	return &Hero{
		id:          id,
		Name:        name,
		Description: description,
		Type:        heroType,
		Role:        role,
		Stats:       statistics,
		ActionList:  actionList,
		Cooldowns:   make([]int, len(actionList)),
		HeroImage:   assets.MustLoadImage(heroImage),
		IconImage:   assets.MustLoadImage(iconImage),
		BannerImage: assets.MustLoadImage(bannerImage),
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
		Type:        uint8(rand.Intn(7)),
		Description: "A hero",
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
