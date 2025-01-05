package stats

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
		Attack:      100,
		P_Def:      100,
		A_Def:      100,
		W_Def:      100,
		Speed:      100,
		ActionDice: 30,
	}
}
