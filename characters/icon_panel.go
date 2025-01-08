package characters

import (
	"quasar/assets"
	"quasar/common"


	"github.com/hajimehoshi/ebiten/v2"
)

var placeholder, _ = assets.MustLoadImage("images/placeholder.png")

type IconPanel struct {
	X, Y       int
	Icons      []*Icon
	horizontal bool
}

type Icon struct {
	X, Y int
	Char Character
}

func NewIconPanel(x, y float32, horizontal bool, length int) *IconPanel {
	ip := &IconPanel{
		X:          int(x),
		Y:          int(y),
		horizontal: horizontal,
	}
	ip.Icons = make([]*Icon, 0)
	for i := 0; i < length; i++ {
		ip.AddIcon(nil)
	}
	return ip
}

func (i *IconPanel) AddIcon(c Character) {
	icon := &Icon{Char: c}
	nextX, nextY := i.X, i.Y+len(i.Icons)*120
	if i.horizontal {
		nextX, nextY = i.X+len(i.Icons)*110, i.Y
	}
	icon.X = nextX
	icon.Y = nextY
	i.Icons = append(i.Icons, icon)
}

func (i *IconPanel) UpdateIcon(c Character, index int) {
	i.Icons[index].Char = c
}

func (i *IconPanel) DrawIcons(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for _, icon := range i.Icons {
		op.GeoM.Reset()
		op.GeoM.Translate(float64(icon.X), float64(icon.Y))
		if icon.Char != nil {
			icon.Char.DrawIcon(screen, icon.X, icon.Y)
		} else {
			screen.DrawImage(placeholder, op)
		}

	}
}

func (i *Icon) GetBounds() (float32, float32, float32, float32) {
	x, y := float32(i.X), float32(i.Y)
	return x, y, x + 100, y + 100
}

func (i *IconPanel) DrawTooltip(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()
	for _, icon := range i.Icons {
		if common.Collide(x, y, icon) {
			icon.Char.DrawTooltip(screen, x, y)
		}
	}
}
