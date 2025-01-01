package assets

import (
	"image"
	"embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed "*"
var assets embed.FS

func MustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
