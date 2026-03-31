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
		g.state.paddle1 += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.state.paddle1 -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.state.paddle2 += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.state.paddle2 -= 5
	}
	g.state.PollState()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	h := screen.Bounds().Dy()
	scale := float64(h) / Height

	drawRect(PaddleWidth*float64(scale), PaddleHeight*float64(scale), 0, g.state.paddle1*float64(scale), screen)
	drawRect(PaddleWidth*float64(scale), PaddleHeight*float64(scale), (Width-PaddleWidth)*float64(scale), g.state.paddle2*float64(scale), screen)
	drawRect(BallWidth*float64(scale), BallHeight*float64(scale), g.state.ball.x*float64(scale), g.state.ball.y*float64(scale), screen)
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

func drawRect(w, h, x, y float64, s *ebiten.Image) {

	newImage := ebiten.NewImage(1, 1)
	newImage.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(w, h)
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(color.White)
	s.DrawImage(newImage, op)
}
