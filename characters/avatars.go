package characters

import (
	"image/color"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type AvatarSquad struct {
	X, Y  float32
	Squad []*Avatar
}

type Avatar struct {
	X, Y          float32
	Width, Height float32
	Image        *ebiten.Image
}

func NewAvatarSquad(x, y float32) *AvatarSquad {
	return &AvatarSquad{X: x, Y: y, Squad: make([]*Avatar, 0)}
}

func NewAvatar(c Character, x, y float32) *Avatar {
	avatar := &Avatar{Image: c.GetAvatar()}
	avatar.X, avatar.Y = x, y
	avatar.Width, avatar.Height = float32(avatar.Image.Bounds().Dx()), float32(avatar.Image.Bounds().Dy())
	return avatar
}

func (s *AvatarSquad) AddAvatar(c Character) {
	avatar := &Avatar{Image: c.GetAvatar()}
	nextX, nextY := s.X+float32(len(s.Squad)*125), s.Y
	avatar.X, avatar.Y = nextX, nextY
	s.Squad = append(s.Squad, avatar)
}

// active is the Team index of the active character, -1 for none
func (s *AvatarSquad) Draw(screen *ebiten.Image, active int) {
	for i, avatar := range s.Squad {
		avatar.Draw(screen, active == i, false)
	}
}

func (a *Avatar) Draw(screen *ebiten.Image, active, targetable bool) {
	var highlightColor *color.RGBA
	if targetable {
		highlightColor = common.RedColor
	} else if active {
		highlightColor = common.GoldColor
	}
	if highlightColor != nil {
		r := max(a.Width / 2, a.Height / 2)
		vector.StrokeCircle(screen, a.X+r, a.Y+r, r, 3, highlightColor, false)

	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(float64(a.X), float64(a.Y))
	screen.DrawImage(a.Image, op)
}

func (a *Avatar) GetBounds() (float32, float32, float32, float32) {
	return a.X, a.Y, a.X + a.Width, a.Y + a.Height
}
