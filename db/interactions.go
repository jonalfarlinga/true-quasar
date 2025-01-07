package db

import (
	"database/sql"
	"fmt"
	"quasar/characters"
	"quasar/characters/actions"
	"quasar/characters/stats"
)

func InsertHero(db *sql.DB, hero *characters.Hero) error {
	query := `INSERT INTO Heroes (
		name, description, type, role, action_list,
		res, Attack, P_Def, A_Def, W_Def,
		P_Boost, A_Boost, W_Boost, Speed, ActionDice
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
    );`
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
		return fmt.Errorf("error inserting into Heroes table: %v", err)
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
		st := stats.NewStats(
			res, attack, pDef, aDef, wDef,
			pBoost, aBoost, wBoost, speed, actionDice,
		)
		al := actions.NewActionList(actionList)
		hero := characters.NewHero(
			id, name, description, uint8(heroType), uint8(role), st, al,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning Heroes table: %v", err)
		}
		heroes = append(heroes, hero)
	}
	return heroes, nil
}
