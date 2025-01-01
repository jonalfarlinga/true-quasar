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

func StatsDefault() *Statistics {
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
