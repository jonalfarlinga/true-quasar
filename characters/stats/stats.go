package stats

import "fmt"

type Statistics struct {
	Resilience int
	ResCurrent int
	Attack     int
	P_Def      int
	A_Def      int
	W_Def      int
	P_Boost    int
	A_Boost    int
	W_Boost    int
	Speed      int
	ActionDice int
}

func NewStats(resilience, Atk, pDef, aDef, wDef, pBoost, aBoost, wBoost, speed, actionDice int) *Statistics {
	return &Statistics{
		Resilience: resilience,
		ResCurrent: resilience,
		Attack:     Atk,
		P_Def:      pDef,
		A_Def:      aDef,
		W_Def:      wDef,
		P_Boost:    pBoost,
		A_Boost:    aBoost,
		W_Boost:    wBoost,
		Speed:      speed,
		ActionDice: actionDice,
	}
}

func DefaultStats() *Statistics {
	return &Statistics{
		Resilience: 50,
		ResCurrent: 50,
		Attack:     100,
		P_Def:      100,
		A_Def:      100,
		W_Def:      100,
		Speed:      100,
		ActionDice: 30,
	}
}

func ZeroStats() *Statistics {
	return &Statistics{
		Resilience: 0,
		ResCurrent: 0,
		Attack:     0,
		P_Def:      0,
		A_Def:      0,
		W_Def:      0,
		P_Boost:    0,
		A_Boost:    0,
		W_Boost:    0,
		Speed:      0,
		ActionDice: 0,
	}
}

func (s *Statistics) ToString() string {
	return fmt.Sprintf("Resilience: %d\nAttack: %d\nPower Def: %d\nAccuracy Def: %d\nWill Def: %d\nPower Boost: %d\nAccuracy Boost: %d\nWill Boost: %d\nSpeed: %d\nAction Dice: %d",
		s.Resilience, s.Attack, s.P_Def, s.A_Def, s.W_Def, s.P_Boost, s.A_Boost, s.W_Boost, s.Speed, s.ActionDice)
}

func (s *Statistics) Copy() *Statistics {
	return &Statistics{
		Resilience: s.Resilience,
		ResCurrent: s.ResCurrent,
		Attack:     s.Attack,
		P_Def:      s.P_Def,
		A_Def:      s.A_Def,
		W_Def:      s.W_Def,
		P_Boost:    s.P_Boost,
		A_Boost:    s.A_Boost,
		W_Boost:    s.W_Boost,
		Speed:      s.Speed,
		ActionDice: s.ActionDice,
	}
}
