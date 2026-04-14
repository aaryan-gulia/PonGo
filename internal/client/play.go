package client

import (
	"PonGo/internal/pong"
	"bytes"
	"encoding/gob"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"net"
)

type Game struct {
	state pong.GameState
}

type GameOnlinePVP struct {
	Game
	conn *net.UDPConn
	end  chan struct{}
}

func (g *GameOnlinePVP) setup() {

	g.state.Reset()

	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Fatal("Couldn’t resolve address:", err)
	}
	g.conn, err = net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	g.end = make(chan struct{})

	go g.pollState()
}

func (g *GameOnlinePVP) close() {
	g.conn.Close()
	close(g.end)
}

func (g *GameOnlinePVP) update() (ClientState, error) {
	e := pollEvent()
	if e == pong.Q {
		return MenuPage, nil
	}
	g.handleEvent(e)
	return GameOnlinePVPPage, nil
}

func pollEvent() pong.GameEvent {
	e := pong.N
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		e = pong.W
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		e = pong.S
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		e = pong.Q
	}
	return e
}

func (g *GameOnlinePVP) handleEvent(e pong.GameEvent) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(e); err != nil {
		log.Fatalln("encoding error : ", err)
	}
	_, err := g.conn.Write(buf.Bytes())
	if err != nil {
		log.Println("writing error : ", err)
	}
}

func (g *GameOnlinePVP) pollState() {
	buffer := make([]byte, 1024)
	for {
		select {
		case <-g.end:
			return
		default:
		}
		n, _, err := g.conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("read error: ", err)
			return
		}
		var s pong.GameState
		dec := gob.NewDecoder(bytes.NewReader(buffer[:n]))
		if err := dec.Decode(&s); err != nil {
			log.Println("decoding error : ", err)
			return
		}
		g.state = s
	}
}

func (g *Game) draw(screen *ebiten.Image) {
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

type GameAI struct {
	Game
}

func (g *GameAI) setup() {
	g.state.Reset()
	log.Println("starting ai game")
}
func (g *GameAI) close() {}

func (g *GameAI) update() (ClientState, error) {
	e := pollEvent()
	if e == pong.Q {
		return MenuPage, nil
	}
	g.handleEvent(e)
	return GameAIPage, nil
}

func (g *GameAI) handleEvent(e pong.GameEvent) {
	if e == pong.W {
		g.state.MovePaddle1Up()
	}
	if e == pong.S {
		g.state.MovePaddle1Down()
	}
	if g.state.Ball.Y > g.state.Paddle2+pong.PaddleHeight/2 {
		g.state.MovePaddle2Down()
	} else {
		g.state.MovePaddle2Up()
	}
	g.state.PollState()
}
