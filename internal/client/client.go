package client

import (
	"PonGo/internal/pong"
	"bytes"
	"encoding/gob"
	"image/color"
	"log"
	"net"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameEvent int
type ConnectionState int
type GameType int
type ClientState int

const (
	W GameEvent = iota
	S
	P
	Q
)

const (
	Connected ConnectionState = iota
	Waiting
	Paused
	Playing
	Disconnected
)

const (
	Ai GameType = iota
	OnlinePvp
)

const (
	MainMenu ClientState = iota
	GamePlay
)

type Game struct {
	state pong.GameState
	conn  *net.UDPConn
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		handleEvent(W, g.conn)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		handleEvent(S, g.conn)
	}

	return nil
}

func handleEvent(e GameEvent, conn *net.UDPConn) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(e); err != nil {
		log.Fatalln("encoding error : ", err)
	}

	_, err := conn.Write(buf.Bytes())
	if err != nil {
		log.Println("writing error : ", err)
	}

}

func (g *Game) Draw(screen *ebiten.Image) {
	h := screen.Bounds().Dy()
	scale := float64(h) / pong.Height

	drawRect(pong.PaddleWidth*float64(scale), pong.PaddleHeight*float64(scale), 0, g.state.Paddle1*float64(scale), screen)
	drawRect(pong.PaddleWidth*float64(scale), pong.PaddleHeight*float64(scale), (pong.Width-pong.PaddleWidth)*float64(scale), g.state.Paddle2*float64(scale), screen)
	drawRect(pong.BallWidth*float64(scale), pong.BallHeight*float64(scale), g.state.Ball.X*float64(scale), g.state.Ball.Y*float64(scale), screen)

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

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return 640, 480
}

func Run() {

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("game client")
	var state pong.GameState

	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Fatal("Couldn’t resolve address:", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	defer conn.Close()

	state.Reset()
	if err := ebiten.RunGame(&Game{state: state, conn: conn}); err != nil {
		log.Fatalln("game engine error :", err)
	}
}
