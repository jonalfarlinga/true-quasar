package assets

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed "*"
var assetLib embed.FS

func MustLoadImage(name string) (*ebiten.Image, error) {
	f, err := assetLib.Open(cleanAddress(name))
	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %v", name, err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("could not decode image %s: %v", name, err)
	}

	return ebiten.NewImageFromImage(img), nil
}

func cleanAddress(address string) string {
	address = strings.ToLower(address)
	address = strings.ReplaceAll(address, " ", "_")
	return address
}
