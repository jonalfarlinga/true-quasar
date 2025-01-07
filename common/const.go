package common

import (
	"image/color"

	"golang.org/x/image/font/basicfont"
)

const (
	StateMenu uint8 = iota
	StateDraft
	StateCombat
	StateCreate
)

const (
	ScreenWidth float32 = 1280.0
	ScreenHeight float32 = 768.0
	FrameRate = 30
)

const (
	Power uint8 = iota
	Accuracy
	Will
)

var (
	BackgroundColor  = color.RGBA{0x10, 0x10, 0x10, 0xff}
	ButtonColor      = color.RGBA{0x50, 0x50, 0x00, 0xff}
	ButtonHoverColor = color.RGBA{0x70, 0x70, 0x30, 0xff}
	ButtonOffColor   = color.RGBA{0x30, 0x30, 0x30, 0xff}
	GoldColor        = color.RGBA{0xff, 0xd7, 0x00, 0xff}
	RedColor         = color.RGBA{0xDD, 0x20, 0x15, 0xff}
	PanelColor       = color.RGBA{0x30, 0x30, 0xA0, 0xFF}
	WhiteColor	   		= color.White
)

// Hero Types
const (
	Nexus uint8 = iota
	Antimatter
	Graviton
	Plasma
	Metal
	Void
)

var (
	HeroTypeNames = map[uint8]string{
		Nexus:      "Nexus",
		Antimatter: "Antimatter",
		Graviton:   "Graviton",
		Plasma:     "Plasma",
		Metal:      "Metal",
		Void:       "Void",
	}
)

// Hero Roles
const (
	Striker uint8 = iota
	Defender
	Controller
	Channeler
)

var (
	HeroRoleNames = map[uint8]string{
		Striker:    "Striker",
		Defender:   "Defender",
		Controller: "Controller",
		Channeler:  "Channeler",
	}
)

var (
	MenuFont = basicfont.Face7x13
)
