package common

import "image/color"

const (
	StateMenu = iota
	StateDraft
	StateCombat
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 768
)

const (
	Power = iota
	Accuracy
	Will
)

var (
	BackgroundColor = color.RGBA{R: 0x10, G: 0x10, B: 0x10, A: 0xff}
)

const (
	Nexus = iota
	Antimatter
	Graviton
	Plasma
	Metal
	Void
)
