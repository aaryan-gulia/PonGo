package pong

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

	vector.DrawFilledRect(screen, float32(g.state.ball.x), float32(g.state.ball.y), float32(BallWidth), float32(BallHeight), color.White, false)
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
