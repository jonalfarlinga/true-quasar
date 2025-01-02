package db

import (
	"database/sql"
	"fmt"
	"quasar/characters"
	"quasar/characters/actions"
	"quasar/characters/stats"
)

func GetAllHeroes(db *sql.DB) ([]*characters.Hero, error) {
	query := `SELECT * FROM Heroes;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying Heroes table: %v", err)
	}
	defer rows.Close()

	heroes := []*characters.Hero{}
	for rows.Next() {
		var (
			name, description, heroImage, iconImage, bannerImage, actionList string
			id, heroType, role, res, pAtk, aAtk, wAtk, pDef, aDef, wDef      int
			pBoost, aBoost, wBoost, speed, actionDice                        int
		)
		err := rows.Scan(
			&id, &name, &description, &heroType, &role, &actionList,
			&heroImage, &iconImage, &bannerImage,
			&res, &pAtk, &aAtk, &wAtk, &pDef, &aDef, &wDef,
			&pBoost, &aBoost, &wBoost, &speed, &actionDice,
		)
		st := stats.NewStats(res, pAtk, aAtk, wAtk, pDef, aDef, wDef, pBoost, aBoost, wBoost, speed, actionDice)
		ac := actions.NewActionList(actionList)
		hero := characters.NewHero(
			id, name, description, uint8(heroType), uint8(role),
			st, ac, heroImage, iconImage, bannerImage,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning Heroes table: %v", err)
		}
		heroes = append(heroes, hero)
	}
	return heroes, nil
}
