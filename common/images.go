package common

import (
	"quasar/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

var Emblems []*ebiten.Image

func LoadEmblems() {
	nexusEmblem, _ := assets.MustLoadImage("images/emblem/nexus.png")
	antimatterEmblem, _ := assets.MustLoadImage("images/emblem/antimatter.png")
	gravitonEmblem, _ := assets.MustLoadImage("images/emblem/graviton.png")
	plasmaEmblem, _ := assets.MustLoadImage("images/emblem/plasma.png")
	metalEmblem, _ := assets.MustLoadImage("images/emblem/metal.png")
	voidEmblem, _ := assets.MustLoadImage("images/emblem/void.png")
	strikerEmblem, _ := assets.MustLoadImage("images/emblem/striker.png")
	defenderEmblem, _ := assets.MustLoadImage("images/emblem/defender.png")
	channelerEmblem, _ := assets.MustLoadImage("images/emblem/channeler.png")
	controllerEmblem, _ := assets.MustLoadImage("images/emblem/controller.png")
	Emblems = []*ebiten.Image{nexusEmblem, antimatterEmblem, gravitonEmblem, plasmaEmblem, metalEmblem, voidEmblem, strikerEmblem, defenderEmblem, channelerEmblem, controllerEmblem}
}
