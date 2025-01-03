package db

import (
	"database/sql"
	"quasar/characters"
	"quasar/characters/actions"
	"quasar/characters/stats"
	"quasar/common"
)

var RepoList = []*characters.Hero{
	// Hero1
	characters.NewHero(
		-1,
		"Vanguard",
		"A hero that excels at taking hits and protecting allies.",
		common.Metal,
		common.Defender,
		stats.NewStats(70, 90, 80, 80, 120, 120, 120, 1, 0, 0, 85, 30),
		actions.NewActionList(
			"Charging Slam,Aegis,Bastion",
		),
	),
	// Hero2
	characters.NewHero(
		-1,
		"Quasar Knight",
		"A hero that excels at controlling the battlefield and protecting allies.",
		common.Nexus,
		common.Defender,
		stats.NewStats(60, 100, 100, 100, 100, 100, 100, 1, 1, 1, 90, 30),
		actions.NewActionList(
			"Void Implosion,Graviton Lock,Aegis",
		),
	),
	// Hero3
	characters.NewHero(
		-1,
		"Goliath",
		"A hero that excels at dealing damage and taking hits.",
		common.Plasma,
		common.Defender,
		stats.NewStats(80, 80, 80, 120, 120, 120, 120, 2, 0, 1, 80, 30),
		actions.NewActionList(
			"Antimatter Strike,Aegis",
		),
	),
	// Hero4
	characters.NewHero(
		-1,
		"Gravitus",
		"A hero that excels at controlling the battlefield and protecting allies.",
		common.Graviton,
		common.Defender,
		stats.NewStats(60, 80, 80, 80, 100, 130, 130, 1, 0, 0, 90, 30),
		actions.NewActionList(
			"Graviton Lock,Deflection Field,Bastion",
		),
	),
	// Hero5
	characters.NewHero(
		-1,
		"Assassin",
		"A hero that excels at dealing damage and taking out key targets.",
		common.Antimatter,
		common.Striker,
		stats.NewStats(45, 130, 130, 100, 90, 100, 80, 0, 2, 0, 95, 30),
		actions.NewActionList(
			"Antimatter Strike,Plasma Blast,Void Snare",
		),
	),
	// Hero6
	characters.NewHero(
		-1,
		"Voidwalker",
		"A hero that excels at dealing damage and controlling the battlefield.",
		common.Void,
		common.Striker,
		stats.NewStats(50, 120, 120, 120, 90, 90, 90, 0, 1, 1, 90, 30),
		actions.NewActionList(
			"Void Implosion,Void Snare,Void Implosion",
		),
	),
	// Hero7
	characters.NewHero(
		-1,
		"Riftblade",
		"A hero that excels at dealing damage and taking out key targets.",
		common.Plasma,
		common.Striker,
		stats.NewStats(45, 130, 130, 100, 90, 100, 80, 0, 4, 0, 75, 30),
		actions.NewActionList(
			"Antimatter Strike,Plasma Blast,Graviton Surge",
		),
	),
	// Hero8
	characters.NewHero(
		-1,
		"Battle Angel",
		"A hero that excels at dealing damage and controlling the battlefield.",
		common.Graviton,
		common.Striker,
		stats.NewStats(50, 120, 120, 120, 90, 90, 90, 0, 1, 1, 110, 30),
		actions.NewActionList(
			"Void Implosion,Deflection Field,Graviton Surge",
		),
	),
	// Hero9
	characters.NewHero(
		-1,
		"Chronomancer",
		"A hero that excels at controlling the battlefield and protecting allies.",
		common.Nexus,
		common.Controller,
		stats.NewStats(60, 100, 100, 150, 75, 75, 75, 1, 1, 3, 105, 30),
		actions.NewActionList(
			"Plasma Barrage,Temporal Rift,Plasma Matrix",
		),
	),
	// Hero10
	characters.NewHero(
		-1,
		"Void Prophet",
		"A hero that excels at controlling the battlefield and protecting allies.",
		common.Void,
		common.Controller,
		stats.NewStats(50, 80, 80, 125, 110, 110, 110, 1, 1, 1, 95, 30),
		actions.NewActionList(
			"Void Implosion,Void Infusion,Cosmic Resonance",
		),
	),
	// Hero11
	characters.NewHero(
		-1,
		"Dark Sorcerer",
		"A hero that excels at dealing damage and controlling the battlefield.",
		common.Antimatter,
		common.Controller,
		stats.NewStats(45, 130, 130, 100, 90, 100, 80, 0, 2, 0, 100, 30),
		actions.NewActionList(
			"Antimatter Strike,Antimatter Cascade,Temporal Rift",
		),
	),
	// Hero12
	characters.NewHero(
		-1,
		"Terrifier",
		"A hero that excels at dealing damage and controlling the battlefield.",
		common.Metal,
		common.Controller,
		stats.NewStats(70, 90, 80, 80, 120, 120, 120, 1, 0, 0, 95, 30),
		actions.NewActionList(
			"Charging Slam,Graviton Collapse,Graviton Surge",
		),
	),
	// Hero13
	characters.NewHero(
		-1,
		"Void Knight",
		"A hero that leads from the front, energizing allies",
		common.Void,
		common.Channeler,
		stats.NewStats(50, 80, 80, 125, 110, 110, 110, 1, 1, 1, 95, 30),
		actions.NewActionList(
			"Void Implosion,Void Infusion,Plasma Overload",
		),
	),
	// Hero14
	characters.NewHero(
		-1,
		"Grand Unifier",
		"A hero who heals and buffs allies",
		common.Plasma,
		common.Channeler,
		stats.NewStats(60, 100, 100, 150, 75, 75, 75, 1, 1, 3, 100, 30),
		actions.NewActionList(
			"Plasma Barrage,Plasma Matrix,Cosmic Resonance",
		),
	),
	// Hero15
	characters.NewHero(
		-1,
		"Warden",
		"A hero that excels at taking hits and protecting allies.",
		common.Metal,
		common.Channeler,
		stats.NewStats(70, 90, 80, 80, 120, 120, 120, 1, 0, 0, 90, 30),
		actions.NewActionList(
			"Charging Slam,Graviton Collapse,Graviton Surge",
		),
	),
	// Hero16
	characters.NewHero(
		-1,
		"Voidbringer",
		"A hero who disrupts enemies and empowers allies",
		common.Graviton,
		common.Channeler,
		stats.NewStats(60, 80, 80, 80, 100, 130, 130, 1, 0, 0, 95, 30),
		actions.NewActionList(
			"Graviton Lock,Deflection Field,Plasma Overload",
		),
	),
}

func installHeroRepo(db *sql.DB) {
	for _, hero := range RepoList {
		InsertHero(db, hero)
	}
}
