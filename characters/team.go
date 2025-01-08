package characters

import (
	"image/color"
	"quasar/characters/actions"
	"quasar/characters/stats"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Team struct {
	Heroes      []*Hero
	DamageDealt []int
	DamageTaken []int
	HealsDealt  []int
}

func NewTeam() *Team {
	t := &Team{}
	t.Heroes = make([]*Hero, 4)
	for i := 0; i < 4; i++ {
		t.Heroes[i] = DefaultHero("vanguard", stats.DefaultStats(), actions.ActionsDefault())
	}
	t.DamageDealt = make([]int, 4)
	t.DamageTaken = make([]int, 4)
	t.HealsDealt = make([]int, 4)
	return t
}

func (t *Team) DrawIcons(screen *ebiten.Image, x, y float64, horiz bool) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	emb_op := &ebiten.DrawImageOptions{}
	emb_op.GeoM.Scale(0.25, 0.25)
	emb_op.GeoM.Translate(x, y)
	for i := 0; i < 4; i++ {
		hero := t.Heroes[i]
		if hero != nil {
			screen.DrawImage(hero.IconImage, op)
			text.Draw(screen, hero.Name, common.MenuFont, int(op.GeoM.Element(0, 2))+10, int(op.GeoM.Element(1, 2))+110, color.White)
			screen.DrawImage(common.Emblems[hero.Type], emb_op)
			emb_op.GeoM.Translate(75, 0)
			screen.DrawImage(common.Emblems[hero.Role+6], emb_op)
			emb_op.GeoM.Translate(-75, 0)
		} else {
			screen.DrawImage(placeholder, op)
		}
		if horiz {
			op.GeoM.Translate(110, 0)
			emb_op.GeoM.Translate(110, 0)

		} else {
			op.GeoM.Translate(0, 120)
			emb_op.GeoM.Translate(0, 120)
		}
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
