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
	BackgroundColor = color.Gray16{0x111F}
	MenuButtonColor = color.Gray16{0x555F}
	MenuButtonHoverColor = color.Gray16{0x7FFF}
)

const (
	Nexus uint8 = iota
	Antimatter
	Graviton
	Plasma
	Metal
	Void
)
