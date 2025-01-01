package characters

import (
	"math/rand"
	"quasar/assets"
	"quasar/characters/actions"
	"quasar/characters/stats"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hero struct {
	Name        string
	Description string
	Type        uint8
	Stats       *stats.Statistics
	ActionList  []*actions.Action
	Cooldowns   []int
	HeroImage   *ebiten.Image
	BannerImage *ebiten.Image
}

func (h *Hero) GetStats() *stats.Statistics {
	return h.Stats
}

func (h *Hero) GetActionList() []*actions.Action {
	return h.ActionList
}

func (h *Hero) DrawHero() {
	// Draw the hero on the screen
}

func (h *Hero) DrawCard() {
	// Draw the hero's card on the screen
}

func NewHero(name string, statistics *stats.Statistics, actionList []*actions.Action) *Hero {
	if statistics == nil {
		statistics = stats.StatsDefault()
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
		HeroImage:   assets.MustLoadImage("assets/images/vanguard_hero.png"),
		BannerImage: assets.MustLoadImage("assets/images/vanguard_banner.png"),
	}
}
