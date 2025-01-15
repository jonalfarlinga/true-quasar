package db

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"
	"quasar/characters"
	"quasar/characters/actions"
	"quasar/characters/stats"
	"quasar/common"
	"strconv"
)

var RepoList = []*characters.Hero{
	// Hero1
	characters.NewHero(
		-1,
		"Vanguard",
		"A front line fighter that excels at taking hits and protecting allies.",
		common.Metal,
		common.Defender,
		stats.NewStats(70, 85, 120, 120, 120, 1, 0, 0, 85, 30),
		actions.NewActionList(
			"Charging Slam,Aegis,Bastion",
		),
	),
	// Hero2
	characters.NewHero(
		-1,
		"Quasar Knight",
		"A legendary duelist who focuses on damaging single targets while reducing their effectiveness.",
		common.Nexus,
		common.Defender,
		stats.NewStats(60, 100, 100, 100, 100, 1, 1, 1, 92, 30),
		actions.NewActionList(
			"Charging Slam,Graviton Surge,Electron Cage",
		),
	),
	// Hero3
	characters.NewHero(
		-1,
		"Goliath",
		"A hero that excels at dealing damage and taking hits.",
		common.Plasma,
		common.Defender,
		stats.NewStats(80, 120, 120, 120, 120, 2, 0, 1, 80, 30),
		actions.NewActionList(
			"Antimatter Strike,Aegis",
		),
	),
	// Hero4
	characters.NewHero(
		-1,
		"Gravus Krang",
		"A tough hero that gains defense as he loses resilience, and takes damage for other heroes.",
		common.Graviton,
		common.Defender,
		stats.NewStats(50, 90, 130, 100, 130, 1, 0, 1, 86, 30),
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
		stats.NewStats(45, 130, 90, 100, 80, 0, 2, 0, 95, 30),
		actions.NewActionList(
			"Antimatter Strike,Plasma Blast,Void Snare",
		),
	),
	// Hero6
	characters.NewHero(
		-1,
		"Void Walker",
		"An enigmatic hero that deals damage over time and attacks multiple enemies.",
		common.Void,
		common.Striker,
		stats.NewStats(50, 120, 95, 85, 85, 0, 0, 1, 91, 30),
		actions.NewActionList(
			"Precision Strike,Void Implosion,Graviton Collapse",
		),
	),
	// Hero7
	characters.NewHero(
		-1,
		"Riftblade",
		"An ancient order of mercenaries that specializes in avoiding attacks while taking down hard targets.",
		common.Plasma,
		common.Striker,
		stats.NewStats(45, 125, 85, 100, 80, 0, 4, 0, 89, 30),
		actions.NewActionList(
			"Antimatter Strike,Plasma Blast,Graviton Surge",
		),
	),
	// Hero8
	characters.NewHero(
		-1,
		"Battle Angel",
		"A free-willed android who joins battles seemingly at random. She is a powerful dealer and preventer of damage.",
		common.Graviton,
		common.Striker,
		stats.NewStats(50, 135, 85, 70, 90, 0, 1, 1, 110, 30),
		actions.NewActionList(
			"Gravitic Riposte,Deflection Field,Graviton Collapse",
		),
	),
	// Hero9
	characters.NewHero(
		-1,
		"Chronomancer",
		"The frontline officer of the Nexus, chronomancers control the battlefield through manipulating the gravitic-temporal flow.",
		common.Nexus,
		common.Controller,
		stats.NewStats(60, 100, 75, 75, 85, 1, 1, 3, 105, 30),
		actions.NewActionList(
			"Graviton Lance,Temporal Rift,Plasma Matrix",
		),
	),
	// Hero10
	characters.NewHero(
		-1,
		"Void Prophet",
		"Bringing the inscrutable will of the Void, the prophet deals damage over time and slows enemies.",
		common.Void,
		common.Controller,
		stats.NewStats(50, 85, 115, 115, 115, 1, 0, 1, 95, 30),
		actions.NewActionList(
			"Concentrated Strike,Void Implosion,Void Snare",
		),
	),
	// Hero11
	characters.NewHero(
		-1,
		"Dark Sorcerer",
		"Dabbling in the forbidden arts, the sorcerer deals damage and disrupts enemies.",
		common.Antimatter,
		common.Controller,
		stats.NewStats(45, 130, 85, 100, 85, 0, 2, 0, 100, 30),
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
		stats.NewStats(70, 90, 120, 120, 120, 1, 0, 0, 95, 30),
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
		stats.NewStats(50, 80, 110, 110, 110, 1, 1, 1, 95, 30),
		actions.NewActionList(
			"Void Implosion,Void Infusion,Plasma Overload",
		),
	),
	// Hero14
	characters.NewHero(
		-1,
		"Grand Unifier",
		"An exemplary leader of men, who deals area damage and hastens allies.",
		common.Plasma,
		common.Channeler,
		stats.NewStats(60, 115, 75, 75, 95, 1, 1, 2, 100, 30),
		actions.NewActionList(
			"Plasma Slash,Temporal Rift,Plasma Overload",
		),
	),
	// Hero15
	characters.NewHero(
		-1,
		"Warden",
		"A nurturing hero that bolsters and heals their allies.",
		common.Metal,
		common.Channeler,
		stats.NewStats(70, 90, 120, 120, 120, 1, 0, 0, 90, 30),
		actions.NewActionList(
			"Charging Slam,Bastion,Cosmic Resonance",
		),
	),
	// Hero16
	characters.NewHero(
		-1,
		"Voidbringer",
		"The voidbringer buffs allies and disrupts enemies, with litte regard for themselves.",
		common.Void,
		common.Channeler,
		stats.NewStats(60, 80, 70, 70, 130, 1, 0, 0, 95, 30),
		actions.NewActionList(
			"Graviton Lock,Deflection Field,Plasma Overload",
		),
	),
}

func installHeroRepo(db *sql.DB) {
	for _, hero := range RepoList {
		InsertOrUpdateHero(db, hero)
	}
}

func loadFromCSV(db *sql.DB) {
	// Load from CSV
	file, err := os.Open("db/heroes.csv")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read record: %s", err)
		}
		types := map[string]uint8{
			"Nexus":      0,
			"Antimatter": 1,
			"Graviton":   2,
			"Plasma":     3,
			"Metal":      4,
			"Void":       5,
		}
		roles := map[string]uint8{
			"Striker":    0,
			"Defender":   1,
			"Controller": 2,
			"Channeler":  3,
		}
		hero := characters.NewHero(
			-1,
			record[0],
			record[13],
			types[record[1]],
			roles[record[2]],
			stats.NewStats(
				parseInt(record[3]),
				parseInt(record[4]),
				parseInt(record[5]),
				parseInt(record[6]),
				parseInt(record[7]),
				parseInt(record[8]),
				parseInt(record[9]),
				parseInt(record[10]),
				parseInt(record[11]),
				30,
			),
			actions.NewActionList(record[12]),
		)

		InsertOrUpdateHero(db, hero)
	}
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to parse int: %s", err)
	}
	return i
}
