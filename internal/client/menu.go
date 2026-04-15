package client

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

var fontFace text.GoTextFace
var mainMenuOptions = [3]string{"Play AI", "Play Online PVP", "Quit"}

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatalln("font loading error : ", err)
	}
	fontFace = text.GoTextFace{Source: s, Size: 24}
}

type MainMenuOption int

const (
	PlayAI MainMenuOption = iota
	PlayOnlinePvp
	Quit
)

type MainMenu struct {
	selected MainMenuOption
}

func (m *MainMenu) setup() {}
func (m *MainMenu) close() {}

func (m *MainMenu) Update() error {
	selected, err := m.update()
	if selected != MenuPage {
		log.Println(selected)
	}
	return err
}

func (m *MainMenu) update() (ClientState, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if m.selected == PlayAI {
			return MenuPage, nil
		}
		m.selected--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if m.selected == Quit {
			return MenuPage, nil
		}
		m.selected++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch m.selected {
		case PlayAI:
			return GameAIPage, nil
		case PlayOnlinePvp:
			return GameOnlinePVPPage, nil
		case Quit:
			return Exit, nil
		}
	}
	return MenuPage, nil
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	m.draw(screen)
}

func (m *MainMenu) draw(screen *ebiten.Image) {
	h := screen.Bounds().Dy()
	w := screen.Bounds().Dx()
	for i, o := range mainMenuOptions {
		if i == int(m.selected) {
			drawText(screen, float64(w)/2, float64(h)/2-50+float64(i)*50, color.RGBA{255, 0, 0, 1}, o)
		}
		drawText(screen, float64(w)/2, float64(h)/2-50+float64(i)*50, color.White, o)
	}
}

func drawText(screen *ebiten.Image, x float64, y float64, c color.Color, t string) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(c)
	text.Draw(screen, t, &fontFace, op)

}

func (m *MainMenu) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func RunMainMenu() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Menu Test")
	if err := ebiten.RunGame(&MainMenu{}); err != nil {
		log.Println("game engine error : ", err)
	}
}
