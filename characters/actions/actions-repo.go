package actions

import "quasar/common"

var list map[string]*Action = map[string]*Action{
	"Boss Hammer": {
		Name:		"Boss Hammer",
		Description: "Deal damage to all the enemies with Power",
		Cooldown:	0,
		Type:		common.Power,
		Fn: func() {
			// Boss Hammer logic goes here
		},
	},
	"Power Strike": {
		Name:        "Power Strike",
		Description: "Deal damage to the enemy with Power",
		Cooldown:    0,
		Type:        common.Power,
		Fn: func() {
			// Attack logic goes here
			// basic power attack
		},
	},
	"Precision Strike": {
		Name:        "Precision Strike",
		Description: "Deal damage to the enemy with Accuracy",
		Cooldown:    0,
		Type:        common.Accuracy,
		Fn: func() {
			// Precision Strike logic goes here
			// basic accuracy attack
		},
	},
	"Concentrated Strike": {
		Name:        "Concentrated Strike",
		Description: "Deal damage to the enemy with Will",
		Cooldown:    0,
		Type:        common.Will,
		Fn: func() {
			// Concentrated Strike logic goes here
			// basic will attack
		},
	},
	"Gravitic Riposte"	: {
		Name:		"Gravitic Riposte",
		Description: "Deal damage to the enemy with Will and shield yourself.",
		Cooldown:	0,
		Type:		common.Will,
		Fn: func() {
			// Gravitic Riposte logic goes here
		},
	},
	"Plasma Barrage": {
		Name:        "Plasma Barrage",
		Description: "Launch 3 plasma bolts, dealing Power damage",
		Cooldown:    0,
		Type:        common.Power,
		Fn: func() {
			// Plasma beam logic goes here
		},
	},
	"Plasma Slash": {
		Name:        "Plasma Slash",
		Description: "Deal Power damage to each enemy.",
		Cooldown:    0,
		Type:        common.Power,
		Fn: func() {
			// Plasma Slash logic goes here
		},
	},
	"Graviton Lance": {
		Name:        "Graviton Lance",
		Description: "Fire a beam that damages and slows enemies using Will",
		Cooldown:    0,
		Type:        common.Will,
		Fn: func() {
			// Graviton Lance logic goes here
		},
	},
	"Antimatter Strike": {
		Name:        "Antimatter Strike",
		Description: "Strike the enemy with antimatter, dealing damage and reducing their defenses using Accuracy",
		Cooldown:    0,
		Type:        common.Accuracy,
		Fn: func() {
			// Antimatter Strike logic goes here
		},
	},
	"Charging Slam": {
		Name:        "Charging Slam",
		Description: "Deal Power damage and stun the target",
		Cooldown:    0,
		Type:        common.Power,
		Fn: func() {
			// Charging Slam logic goes here
		},
	},
	"Void Implosion": {
		Name:        "Void Implosion",
		Description: "Causes damage over time to all enemies using Will",
		Cooldown:    1,
		Type:        common.Will,
		Fn: func() {
			// Void Implosion logic goes here
		},
	},
    "Marked for Death": {
        Name:        "Marked for Death",
        Description: "Mark an enemy, increasing Accuracy boost against them",
        Cooldown:    1,
        Type:        common.Accuracy,
        Fn: func() {
            // Marked for Death logic goes here
        },
    },
    "Quantum Block": {
        Name:        "Quantum Block",
        Description: "When target ally is attacked, this character takes damage instead",
        Cooldown:    2,
        Type:        common.Will,
        Fn: func() {
            // Quantum Block logic goes here
        },
    },
	"Aegis": {
		Name:        "Aegis",
		Description: "Shield yourself and allies from incoming damage",
		Cooldown:    3,
		Type:        common.Power,
		Fn: func() {
			// Aegis logic goes here
		},
	},
    "Targeting Matrix" : {
        Name:       "Targeting Matrix",
        Description: "Increase Accuracy and Will for all allies",
        Cooldown:   3,
        Type:       common.Accuracy,
        Fn: func() {
            // Targeting Matrix logic goes here
        },
    },
	"Graviton Lock": {
		Name:        "Graviton Lock",
		Description: "Stuns all enemies using Will",
		Cooldown:    3,
		Type:        common.Will,
		Fn: func() {
			// Graviton Lock logic goes here
		},
	},
	"Deflection Field": {
		Name:        "Deflection Field",
		Description: "Reflects damage back at the attacker",
		Cooldown:    3,
		Type:        common.Accuracy,
		Fn: func() {
			// Deflection Field logic goes here
		},
	},
	"Plasma Blast": {
		Name:        "Plasma Blast",
		Description: "Deal damage to all enemies using Power, and deals damage over time",
		Cooldown:    3,
		Type:        common.Power,
		Fn: func() {
			// Plasma Blast logic goes here
		},
	},
	"Bastion": {
		Name:        "Bastion",
		Description: "Heal yourself and one ally",
		Cooldown:    3,
		Type:        common.Will,
		Fn: func() {
			// Bastion logic goes here
		},
	},
	"Temporal Rift": {
		Name:        "Temporal Rift",
		Description: "Hastens all allies and slows all enemies using Will",
		Cooldown:    5,
		Type:        common.Will,
		Fn: func() {
			// Temporal Rift logic goes here
		},
	},
	"Void Snare": {
		Name:        "Void Snare",
		Description: "Reduces an enemy's damage and speed",
		Cooldown:    4,
		Type:        common.Accuracy,
		Fn: func() {
			// Void Snare logic goes here
		},
	},
	"Graviton Collapse": {
		Name:        "Graviton Collapse",
		Description: "Deal damage to all enemies using Power and reduce their Power and Accuracy defenses",
		Cooldown:    5,
		Type:        common.Power,
		Fn: func() {
			// Graviton Collapse logic goes here
		},
	},
	"Electron Cage": {
		Name:		"Electron Cage",
		Description: "Deal Accuracy Damage to a single enemy and Stun them",
		Cooldown:	3,
		Type:		common.Accuracy,
		Fn: func() {
			// Electron Cage logic goes here
		},
	},
	"Antimatter Cascade": {
		Name:        "Antimatter Cascade",
		Description: "Deal damage to all enemies using Accuracy and reduce enemy Accuracy boost",
		Cooldown:    3,
		Type:        common.Accuracy,
		Fn: func() {
			// Antimatter Cascade logic goes here
		},
	},
	"Plasma Matrix": {
		Name:        "Plasma Matrix",
		Description: "Slows all enemies and boost allies Power",
		Cooldown:    4,
		Type:        common.Will,
		Fn: func() {
			// Plasma Matrix logic goes here
		},
	},
	"Cosmic Resonance": {
		Name:        "Cosmic Resonance",
		Description: "Heal all allies and boost their Will",
		Cooldown:    3,
		Type:        common.Will,
		Fn: func() {
			// Cosmic Resonance logic goes here
		},
	},
	"Void Infusion": {
		Name:        "Void Infusion",
		Description: "Greatly boost an ally's Power and Will",
		Cooldown:    3,
		Type:        common.Will,
		Fn: func() {
			// Void Infusion logic goes here
		},
	},
	"Graviton Surge": {
		Name:        "Graviton Surge",
		Description: "Deal damage to an enemy using Will and reduce their Speed",
		Cooldown:    3,
		Type:        common.Will,
		Fn: func() {
			// Graviton Surge logic goes here
		},
	},
	"Plasma Overload": {

		Name:        "Plasma Overload",
		Description: "Deal damage using Power to all enemies and heal all allies",
		Cooldown:    4,
		Type:        common.Power,
		Fn: func() {
			// Plasma Overload logic goes here
		},
	},
}
