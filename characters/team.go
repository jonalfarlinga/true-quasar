package characters

import (
	"image/color"
	"quasar/assets"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var placeholder, _ = assets.MustLoadImage("images/placeholder.png")

type Team struct {
	Heroes      []*Hero
	DamageDealt []int
	DamageTaken []int
	HealsDealt  []int
}

func NewTeam() *Team {
	t := &Team{}
	t.Heroes = make([]*Hero, 4)
	t.DamageDealt = make([]int, 4)
	t.DamageTaken = make([]int, 4)
	t.HealsDealt = make([]int, 4)
	return t
}

func (t *Team) DrawIcons(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(common.ScreenWidth)/2+x, y)
	for i := 0; i < 4; i++ {
		hero := t.Heroes[i]
		if hero != nil {
			screen.DrawImage(hero.IconImage, op)
			text.Draw(screen, hero.Name, common.MenuFont, int(op.GeoM.Element(0, 2))+10, int(op.GeoM.Element(1, 2))+160, color.White)
		} else {
			screen.DrawImage(placeholder, op)
		}
		op.GeoM.Translate(160, 0)
	}
}

func (t *Team) CalculateBonuses() {
	antimatter := false
	graviton := false
	plasma := false
	metal := false
	void := false

	for _, hero := range t.Heroes {
		switch hero.Type {
		case common.Antimatter:
			antimatter = true
		case common.Graviton:
			graviton = true
		case common.Plasma:
			plasma = true
		case common.Metal:
			metal = true
		case common.Void:
			void = true
		}
	}

	for _, hero := range t.Heroes {
		switch hero.Type {
		case common.Antimatter:
			s := hero.GetStats()
			if plasma {
				s.A_Boost += 1
			}
			if metal {
				s.P_Def += 10
				s.A_Def += 10
				s.W_Def += 10
			}
		case common.Graviton:
			s := hero.GetStats()
			if metal {
				s.P_Def += 10
				s.A_Def += 10
				s.W_Def += 10
			}
			if void {
				s.W_Boost += 1
			}
		case common.Plasma:
			s := hero.GetStats()
			if antimatter {
				s.ActionDice += 2
			}
			if void {
				s.W_Boost += 1
			}
		case common.Metal:
			s := hero.GetStats()
			if antimatter {
				s.ActionDice += 2
			}
			if graviton {
				s.P_Boost += 1
			}
		case common.Void:
			s := hero.GetStats()
			if graviton {
				s.P_Boost += 1
			}
			if plasma {
				s.A_Boost += 1
			}
		case common.Nexus:
			s := hero.GetStats()
			if antimatter {
				s.ActionDice += 2
			}
			if graviton {
				s.P_Boost += 1
			}
			if plasma {
				s.A_Boost += 1
			}
			if metal {
				s.P_Def += 10
				s.A_Def += 10
				s.W_Def += 10
			}
			if void {
				s.W_Boost += 1
			}
		}
	}
}
