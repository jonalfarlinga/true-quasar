package actions

type Action struct {
	Name string
	Description string
	Cooldown int
	Fn func()
}

func ActionsDefault() []*Action {
	return []*Action{
		{
			Name: "Attack",
			Description: "Deal damage to the enemy",
			Cooldown: 0,
			Fn: func() {
				// Attack logic goes here
			},
		},
	}
}
