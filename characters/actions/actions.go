package actions

import "strings"

type Action struct {
	Name        string
	Description string
	Cooldown    int
	Fn          func()
}

func ActionsDefault() []*Action {
	return []*Action{
		list["Basic Attack"],
	}
}

func NewActionList(actionList string) []*Action {
	actionNames := strings.Split(actionList, ",")
	actions := make([]*Action, 0, len(actionNames))

	for _, name := range actionNames {
		actions = append(actions, list[name])
	}
	return actions
}

func ActionListToString(acts []*Action) string {
	names := make([]string, 0, len(acts))
	for _, act := range acts {
		names = append(names, act.Name)
	}
	return strings.Join(names, ",")
}
