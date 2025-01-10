package common

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawToolTip(screen *ebiten.Image, title string, content string) {
	// draw tooltip
	newLines := 1
	for _, c := range content {
		if c == '\n' {
			newLines++
		}
	}
	var x, y float32 = ScreenWidth/2-100, ScreenHeight/2-250
	var width, height float32 = 300, 100 + 10*float32(newLines)
	vector.DrawFilledRect(screen, x, y, width, height, ButtonOffColor, false)
	text.Draw(screen, title, MenuFont, int(x+10), int(y+20), WhiteColor)
	text.Draw(screen, content, MenuFont, int(x+10), int(y+30), WhiteColor)
}
