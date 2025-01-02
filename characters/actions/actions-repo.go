package actions

var list map[string]*Action = map[string]*Action{
	"Basic Attack": {
		Name:        "Basic Attack",
		Description: "Deal damage to the enemy",
		Cooldown:    0,
		Fn: func() {
			// Attack logic goes here
		},
	},
	"Plasma Beam": {
		Name:        "Plasma Barrage",
		Description: "Launch a series of plasma bolts, dealing high damage to a single target",
		Cooldown: 0,
		Fn: func() {
			// Plasma beam logic goes here
		},
	},
	"Graviton Lance": {
		Name:        "Graviton Lance",
		Description: "Fire a beam that damages and slows enemies",
		Cooldown: 0,
		Fn: func() {
			// Graviton Lance logic goes here
		},
	},
	"Antimatter Strike": {
		Name:        "Antimatter Strike",
		Description: "Strike the enemy with antimatter, dealing damage and reducing their defenses",
		Cooldown: 0,
		Fn: func() {
			// Antimatter Strike logic goes here
		},
	},
	"Charging Slam": {
		Name:        "Charging Slam",
		Description: "Deal damage and stunning the target",
		Cooldown: 0,
		Fn: func() {
			// Charging Slam logic goes here
		},
	},
	"Void Implosion": {
		Name:        "Void Implosion",
		Description: "Causes damage over time to all enemies",
		Cooldown: 0,
		Fn: func() {
			// Void Implosion logic goes here
		},
	},
	"Aegis": {
		Name:        "Aegis",
		Description: "Shield yourself and allies from incoming damage",
		Cooldown: 3,
		Fn: func() {
			// Aegis logic goes here
		},
	},
	"Graviton Lock": {
		Name:        "Graviton Lock",
		Description: "Stuns all enemies",
		Cooldown: 3,
		Fn: func() {
			// Graviton Lock logic goes here
		},
	},
	"Deflection Field": {
		Name:        "Deflection Field",
		Description: "Reflects damage back at the attacker",
		Cooldown: 3,
		Fn: func() {
			// Deflection Field logic goes here
		},
	},
	"Plasma Blast": {
		Name:        "Plasma Blast",
		Description: "Deal damage to all enemies",
		Cooldown: 3,
		Fn: func() {
			// Plasma Blast logic goes here
		},
	},
	"Bastion": {
		Name:        "Bastion",
		Description: "Heal yourself and allies",
		Cooldown: 3,
		Fn: func() {
			// Bastion logic goes here
		},
	},
	"Temporal Rift": {
		Name:        "Temporal Rift",
		Description: "Hastens all allies and slows all enemies",
		Cooldown: 5,
		Fn: func() {
			// Temporal Rift logic goes here
		},
	},
	"Void Snare": {
		Name:        "Void Snare",
		Description: "Reduces enemy damage and speed",
		Cooldown: 4,
		Fn: func() {
			// Void Snare logic goes here
		},
	},
	"Graviton Collapse": {
		Name:        "Graviton Collapse",
		Description: "Deal damage to all enemies and reduce Power and Accuracy defenses",
		Cooldown: 5,
		Fn: func() {
			// Graviton Collapse logic goes here
		},
	},
	"Antimatter Cascade": {
		Name:        "Antimatter Cascade",
		Description: "Deal damage to all enemies and reduce enemy Accuracy boost",
		Cooldown: 3,
		Fn: func() {
			// Antimatter Cascade logic goes here
		},
	},
	"Plasma Matrix": {
		Name:        "Plasma Matrix",
		Description: "Slows all enemies and boost allies Will",
		Cooldown: 4,
		Fn: func() {
			// Plasma Matrix logic goes here
		},
	},
	"Cosmic Resonance": {
		Name:        "Cosmic Resonance",
		Description: "Heal all allies and boost their Will",
		Cooldown: 3,
		Fn: func() {
			// Cosmic Resonance logic goes here
		},
	},
	"Void Infusion": {
		Name:        "Void Infusion",
		Description: "Greatly boost an ally's Power and Will",
		Cooldown: 3,
		Fn: func() {
			// Void Infusion logic goes here
		},
	},
	"Graviton Surge": {
		Name:        "Graviton Surge",
		Description: "Deal damage to an enemy and reduce their Speed",
		Cooldown: 3,
		Fn: func() {
			// Graviton Surge logic goes here
		},
	},
	"Plasma Overload": {

		Name:        "Plasma Overload",
		Description: "Deal damage to all enemies and heal all allies",
		Cooldown: 4,
		Fn: func() {
			// Plasma Overload logic goes here
		},
	},
}
