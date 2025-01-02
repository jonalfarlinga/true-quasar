package draft

import (
	"image/color"
	"quasar/assets"
	"quasar/characters"
	"quasar/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var panelColor = color.RGBA{0x30, 0x30, 0xA0, 0xFF}
var placeholder = assets.MustLoadImage("images/placeholder.png")
var prevMousePressed bool = true

func Update(team *characters.Team) {
	if draftState == DraftStateStart {
		draftLineUp = make([]*DraftCard, 4)
		getDraftPool()
		draftState = DraftStateFirst
		return
	}

	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if mousePressed && !prevMousePressed {
		x, y := ebiten.CursorPosition()
		if common.Collide(x, y, &exitButton) {
			draftState = DraftStateStart
			team.Heroes = make([]*characters.Hero, 4)
			common.GameState = common.StateMenu
		} else if lockInButton.Active && common.Collide(x, y, &lockInButton) {
			if draftState < DraftStateDone {
				lockIn(team)
			}
		} else if draftState == DraftStateDone && common.Collide(x, y, &fightButton) {
			common.GameState = common.StateCombat
		}

		for i, card := range draftLineUp {
			if card.Active && common.Collide(x, y, card) {
				if draftSelection == i {
					draftSelection = -1
				} else if saveSelection == i {
					saveSelection = -1
				} else if draftSelection == -1 {
					draftSelection = i
				} else if i < 4 {
					saveSelection = i
				}

				if draftSelection != -1 && saveSelection != -1 {
					lockInButton.Active = true
				} else {
					lockInButton.Active = false
				}
				break
			}
		}
	}
	prevMousePressed = mousePressed
}

func Draw(screen *ebiten.Image, team *characters.Team) {
	screen.Fill(common.BackgroundColor)
	vector.DrawFilledRect(screen, 10, 10, common.ScreenWidth-20, 180, panelColor, false)
	vector.DrawFilledRect(screen, 10, 210, common.ScreenWidth-20, common.ScreenHeight-220, panelColor, false)

	// Draw team
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(common.ScreenWidth)/2-315, 25)
	for i := 0; i < 4; i++ {
		hero := team.Heroes[i]
		if hero != nil {
			screen.DrawImage(hero.IconImage, op)
			text.Draw(screen, hero.Name, common.MenuFont, int(op.GeoM.Element(0, 2)), int(op.GeoM.Element(1, 2))+158, color.White)
		} else {
			screen.DrawImage(placeholder, op)
		}
		op.GeoM.Translate(160, 0)
	}

	// Draw hero pool
	for i, card := range draftLineUp {
		card.Draw(screen, i == draftSelection, i == saveSelection)
	}

	// Draw buttons
	exitButton.Draw(screen)
	if draftState == DraftStateDone {
		fightButton.Draw(screen)
	} else {
		lockInButton.Draw(screen)
	}
}

func lockIn(team *characters.Team) {
	team.Heroes[draftState] = draftLineUp[draftSelection].Hero
	if draftSelection > 3 {
		draftLineUp = append(draftLineUp[:draftSelection], draftLineUp[draftSelection+1:]...)
	}
	if draftState < DraftStateDone {
		draftLineUp = append(draftLineUp, draftLineUp[saveSelection])
		getDraftPool()
	}
	draftSelection = -1
	saveSelection = -1
	lockInButton.Active = false
	draftState++
	if draftState == DraftStateDone {
		for _, card := range draftLineUp {
			card.Active = false
		}
	}
}
