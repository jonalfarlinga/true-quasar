package common

import "image/color"

const (
	StateMenu uint8 = iota
	StateDraft
	StateCombat
	StateCreate
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 768
)

const (
	Power uint8 = iota
	Accuracy
	Will
)

var (
	BackgroundColor = color.RGBA{0x10, 0x10, 0x10, 0xff}
	MenuButtonColor = color.RGBA{0x50, 0x50, 0x50, 0xff}
	MenuButtonHoverColor = color.RGBA{0x70, 0x70, 0x70, 0xff}
	GoldColor = color.RGBA{0xff, 0xd7, 0x00, 0xff}
	RedColor = color.RGBA{0xDD, 0x20, 0x15, 0xff}
)

const (
	Nexus uint8 = iota
	Antimatter
	Graviton
	Plasma
	Metal
	Void
)
