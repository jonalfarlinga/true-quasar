package common

import (
	"quasar/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

var Emblems []*ebiten.Image

func LoadEmblems() {
	NexusEmblem, _ := assets.MustLoadImage("images/emblem/nexus.png")
	AntimatterEmblem, _ := assets.MustLoadImage("images/emblem/antimatter.png")
	GravitonEmblem, _ := assets.MustLoadImage("images/emblem/graviton.png")
	PlasmaEmblem, _ := assets.MustLoadImage("images/emblem/plasma.png")
	MetalEmblem, _ := assets.MustLoadImage("images/emblem/metal.png")
	VoidEmblem, _ := assets.MustLoadImage("images/emblem/void.png")
	StrikerEmblem, _ := assets.MustLoadImage("images/emblem/striker.png")
	DefenderEmblem, _ := assets.MustLoadImage("images/emblem/defender.png")
	ChannelerEmblem, _ := assets.MustLoadImage("images/emblem/channeler.png")
	ControllerEmblem, _ := assets.MustLoadImage("images/emblem/controller.png")
	Emblems = []*ebiten.Image{AntimatterEmblem, GravitonEmblem, PlasmaEmblem, MetalEmblem, VoidEmblem, NexusEmblem, StrikerEmblem, DefenderEmblem, ChannelerEmblem, ControllerEmblem}
}
