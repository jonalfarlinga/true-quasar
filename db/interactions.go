package db

import (
	"database/sql"
	"fmt"
	"quasar/characters"
	"quasar/characters/actions"
	"quasar/characters/stats"
)

func InsertOrUpdateHero(db *sql.DB, hero *characters.Hero) error {
	query := `INSERT INTO Heroes (
		name, description, type, role, action_list,
		res, Attack, P_Def, A_Def, W_Def,
		P_Boost, A_Boost, W_Boost, Speed, ActionDice
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	) ON DUPLICATE KEY UPDATE
		name = VALUES(name),
		description = VALUES(description),
		type = VALUES(type),
		role = VALUES(role),
		action_list = VALUES(action_list),
		res = VALUES(res),
		Attack = VALUES(Attack),
		P_Def = VALUES(P_Def),
		A_Def = VALUES(A_Def),
		W_Def = VALUES(W_Def),
		P_Boost = VALUES(P_Boost),
		A_Boost = VALUES(A_Boost),
		W_Boost = VALUES(W_Boost),
		Speed = VALUES(Speed),
		ActionDice = VALUES(ActionDice)
	;`
	actionListStr := actions.ActionListToString(hero.GetActionList())
	_, err := db.Exec(
		query,
		hero.Name, hero.Description, hero.Type, hero.Role, actionListStr,
		hero.Stats.Resilience, hero.Stats.Attack,
		hero.Stats.P_Def, hero.Stats.A_Def, hero.Stats.W_Def,
		hero.Stats.P_Boost, hero.Stats.A_Boost, hero.Stats.W_Boost,
		hero.Stats.Speed, hero.Stats.ActionDice,
	)
	if err != nil {
		return fmt.Errorf("error inserting or updating Heroes table: %v", err)
	}

	return nil
}

func GetAllHeroes(db *sql.DB) ([]*characters.Hero, error) {
	query := `SELECT * FROM Heroes;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying Heroes table: %v", err)
	}
	defer rows.Close()
	return parseHeroes(rows)
}

func GetHeroesByType(db *sql.DB, heroType uint8) ([]*characters.Hero, error) {
	query := `SELECT * FROM Heroes WHERE type = ?;`
	rows, err := db.Query(query, heroType)
	if err != nil {
		return nil, fmt.Errorf("error querying Heroes table: %v", err)
	}
	defer rows.Close()
	return parseHeroes(rows)
}

func GetHeroesByRole(db *sql.DB, role uint8) ([]*characters.Hero, error) {
	query := `SELECT * FROM Heroes WHERE role = ?;`
	rows, err := db.Query(query, role)
	if err != nil {
		return nil, fmt.Errorf("error querying Heroes table: %v", err)
	}
	defer rows.Close()
	return parseHeroes(rows)
}

func GetNHeroes(db *sql.DB, n int, by_role, by_type bool, filter uint8) ([]*characters.Hero, error) {
	filter_string := ""
	if by_role {
		filter_string = fmt.Sprintf("WHERE role = %d", filter)
	} else if by_type {
		filter_string = fmt.Sprintf("WHERE type = %d", filter)
	}
	randomizer := fmt.Sprintf(
		"JOIN (SELECT id FROM Heroes %s ORDER BY RAND() LIMIT %d) AS h2",
		filter_string, n,
	)
	query := fmt.Sprintf(`
		SELECT h1.id, h1.name, h1.description, h1.type, h1.role, h1.action_list,
        h1.res, h1.Attack, h1.P_Def, h1.A_Def, h1.W_Def,
        h1.P_Boost, h1.A_Boost, h1.W_Boost, h1.Speed, h1.ActionDice
        FROM Heroes AS h1 %s ON h1.id = h2.id;`,
		randomizer,
	)
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying Heroes table: %v", err)
	}
	defer rows.Close()
	return parseHeroes(rows)
}

func parseHeroes(rows *sql.Rows) ([]*characters.Hero, error) {
	heroes := []*characters.Hero{}
	for rows.Next() {
		var (
			name, description, actionList                     string
			id, heroType, role, res, attack, pDef, aDef, wDef int
			pBoost, aBoost, wBoost, speed, actionDice         int
		)
		err := rows.Scan(
			&id, &name, &description, &heroType, &role, &actionList,
			&res, &attack, &pDef, &aDef, &wDef,
			&pBoost, &aBoost, &wBoost, &speed, &actionDice,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning Heroes table: %v", err)
		}
		st := stats.NewStats(
			res, attack, pDef, aDef, wDef,
			pBoost, aBoost, wBoost, speed, actionDice,
		)
		al := actions.NewActionList(actionList)
		hero := characters.NewHero(
			id, name, description, uint8(heroType), uint8(role), st, al,
		)
		heroes = append(heroes, hero)
	}
	return heroes, nil
}
