package pong

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func Hello() {
	fmt.Println("Hello World")
}

type Game struct {
	state GameState
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		fmt.Println("w")
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		fmt.Println("s")
	}
	g.state.PollState()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ww, wh := ebiten.WindowSize()
	fmt.Println(ww)
	newImage := ebiten.NewImage(1, 1)
	newImage.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(BallWidth*float64(ww)/100, BallHeight*float64(wh)/100)
	op.GeoM.Translate(g.state.ball.x*float64(ww)/100, g.state.ball.y*float64(wh)/100)
	op.ColorScale.ScaleWithColor(color.White)
	screen.DrawImage(newImage, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func Run() {
	var state GameState
	state.Reset()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{state: state}); err != nil {
		fmt.Println("Hello World")
	}
}
