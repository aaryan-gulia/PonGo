package client

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MainMenuOptions int

const (
	PlayAI MainMenuOptions = iota
	PlayOnlinePvp
	Quit
	None
)

type MainMenu struct {
	selected MainMenuOptions
}

func (m *MainMenu) Update() error {
	selected, err := m.update()
	if selected != None {
		log.Println(selected)
	}
	return err
}

func (m *MainMenu) update() (MainMenuOptions, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if m.selected == PlayAI {
			return None, nil
		}
		m.selected--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if m.selected == Quit {
			return None, nil
		}
		m.selected++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return m.selected, nil
	}
	return None, nil
}

func (m *MainMenu) Draw(screen *ebiten.Image) {}

func (m *MainMenu) draw(screen *ebiten.Image) {

}

func (m *MainMenu) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func RunMainMenu() {
	if err := ebiten.RunGame(&MainMenu{}); err != nil {
		log.Println("game engine error : ", err)
	}
}
