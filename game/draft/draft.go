package draft

import (
	"quasar/characters"
	"quasar/common"
	cd "quasar/game/combat/combatdata"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var prevMousePressed bool = true

func Update() {
	if draftState == DraftStateStart {
		cd.CD.Team.Heroes = make([]*characters.Hero, 4)
		draftLineUp = make([]*DraftCard, 4)
		// drafted = make([]*DraftCard, 4)
		updateDraftPool()
		draftState = DraftStateFirst
		return
	}

	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if mousePressed && !prevMousePressed {
		x, y := ebiten.CursorPosition()
		if common.Collide(x, y, &exitButton) {
			draftState = DraftStateStart
			cd.CD.Team.Heroes = make([]*characters.Hero, 4)
			common.GameState = common.StateMenu
		} else if lockInButton.Active && common.Collide(x, y, &lockInButton) {
			if draftState < DraftStateDone {
				lockIn()
			}
		} else if fightButton.Active && common.Collide(x, y, &fightButton) {
			common.GameState = common.StateCombat
			draftState = DraftStateStart
		}

		for i, card := range draftLineUp {
			if card.Active && common.Collide(x, y, card) {
				if draftSelection == i {
					draftSelection = -1
				} else if saveSelection == i {
					saveSelection = -1
				} else if draftSelection == -1 {
					draftSelection = i
				} else if i < 4 && draftState < DraftStateFourth {
					saveSelection = i
				}
				if draftSelection != -1 && (saveSelection != -1 || draftState == DraftStateFourth) {
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

func Draw(screen *ebiten.Image) {
	screen.Fill(common.BackgroundColor)
	vector.DrawFilledRect(screen, 10, 10, common.ScreenWidth-20, 180, common.PanelColor, false)
	vector.DrawFilledRect(screen, 10, 210, common.ScreenWidth-20, common.ScreenHeight-220, common.PanelColor, false)

	// Draw team
	cd.CD.Team.DrawIcons(screen, float64(common.ScreenWidth)/2-215, 70, true)

	// Draw hero pool
	for i, card := range draftLineUp {
		card.Draw(screen, i == draftSelection, i == saveSelection || i > 3)
	}

	// Draw buttons
	exitButton.Draw(screen)
	if draftState == DraftStateDone {
		fightButton.Draw(screen)
	} else {
		lockInButton.Draw(screen)
	}
}

func lockIn() {
	cd.CD.Team.Heroes[draftState] = draftLineUp[draftSelection].Hero
	if draftSelection > 3 {
		draftLineUp = append(draftLineUp[:draftSelection], draftLineUp[draftSelection+1:]...)
	}
	if draftState < DraftStateFourth {
		draftLineUp = append(draftLineUp, draftLineUp[saveSelection])
	}
	updateDraftPool()
	draftSelection = -1
	saveSelection = -1
	lockInButton.Active = false
	draftState++
	if draftState == DraftStateDone {
		fightButton.Active = true
		for _, card := range draftLineUp {
			card.Active = false
		}
	}
}
