package characters

import (
	"quasar/common"
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
	t.DamageDealt = make([]int, 4)
	t.DamageTaken = make([]int, 4)
	t.HealsDealt = make([]int, 4)
	return t
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
				s.A_Atk += 1
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
				s.W_Atk += 1
			}
		case common.Plasma:
			s := hero.GetStats()
			if antimatter {
				s.ActionDice += 2
			}
			if void {
				s.W_Atk += 1
			}
		case common.Metal:
			s := hero.GetStats()
			if antimatter {
				s.ActionDice += 2
			}
			if graviton {
				s.P_Atk += 1
			}
		case common.Void:
			s := hero.GetStats()
			if graviton {
				s.P_Atk += 1
			}
			if plasma {
				s.A_Atk += 1
			}
		case common.Nexus:
			s := hero.GetStats()
			if antimatter {
				s.ActionDice += 2
			}
			if graviton {
				s.P_Atk += 1
			}
			if plasma {
				s.A_Atk += 1
			}
			if metal {
				s.P_Def += 10
				s.A_Def += 10
				s.W_Def += 10
			}
			if void {
				s.W_Atk += 1
			}
		}
	}
}
