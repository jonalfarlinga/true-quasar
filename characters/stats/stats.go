package stats

type Statistics struct {
	Resilience int
	ResCurrent int
	P_Atk      int
	A_Atk      int
	W_Atk      int
	P_Def      int
	A_Def      int
	W_Def      int
	P_Boost    int
	A_Boost    int
	W_Boost    int
	Speed      int
	ActionDice int
}

func NewStats(resilience int, pAtk int, aAtk int, wAtk int, pDef int, aDef int, wDef int, pBoost int, aBoost int, wBoost int, speed int, actionDice int) *Statistics {
	return &Statistics{
		Resilience: resilience,
		ResCurrent: resilience,
		P_Atk:      pAtk,
		A_Atk:      aAtk,
		W_Atk:      wAtk,
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
		P_Atk:      100,
		A_Atk:      100,
		W_Atk:      100,
		P_Def:      100,
		A_Def:      100,
		W_Def:      100,
		Speed:      100,
		ActionDice: 30,
	}
}
