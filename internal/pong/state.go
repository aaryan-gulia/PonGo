package pong

import (
	"log"
	"net"
	"strconv"
	"time"
)

const (
	Width                  float64 = 100
	Height                 float64 = 100
	PaddleHeight           float64 = 16
	PaddleWidth            float64 = 2
	PaddleVelocity         float64 = 3
	BallHeight             float64 = 2
	BallWidth              float64 = 2
	BallVelocityBase       float64 = 1
	BallVelocityMultiplier float64 = 2
)

type GameEvent int

const (
	W GameEvent = iota
	S
	Up
	Down
)

type Ball struct {
	x  float64
	y  float64
	vx float64
	vy float64
	X  float64
	Y  float64
}

type GameState struct {
	Paddle1 float64
	Paddle2 float64
	points1 int
	points2 int
	ball    Ball
	Ball    Ball
}

func (g *GameState) PollState() {
	g.paddleCollision()
	g.wallCollision()
	g.moveBall()

}

func (g *GameState) HandleEvent(conn *net.UDPConn, e GameEvent) {
	switch e {
	case W:
		sendMessage(conn, "w")
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Receive error: %v", err)
			return
		}
		s, _ := strconv.Atoi(string(buffer[:n]))
		log.Println(s)
		g.Paddle1 = float64(s)
	case S:
		sendMessage(conn, "s")
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Receive error: %v", err)
			return
		}
		s, _ := strconv.Atoi(string(buffer[:n]))
		log.Println(s)
		g.Paddle1 = float64(s)
	case Up:
		g.movePaddle2Up()
	case Down:
		g.movePaddle2Down()
	}
}

func sendMessage(conn *net.UDPConn, s string) {
	conn.Write([]byte(s))
}

func (g *GameState) Reset() {
	g.ball.x = Width / 2
	g.ball.y = Height / 2
	g.ball.vx = BallVelocityBase
	g.ball.vy = BallVelocityBase
	g.ball.X = g.ball.x
	g.ball.Y = g.ball.y
	g.Ball = g.ball
}

func (g *GameState) moveBall() {
	g.ball.x += g.ball.vx
	g.ball.X = g.ball.x
	g.ball.y += g.ball.vy
	g.ball.Y = g.ball.y
	g.Ball = g.ball

}

func (g *GameState) paddleCollision() {
	if g.ball.x < PaddleWidth && g.ball.y < g.Paddle1+PaddleHeight && g.ball.y > g.Paddle1 {
		g.ball.vx *= -1
	}
	if g.ball.x > Width-PaddleWidth && g.ball.y < g.Paddle2+PaddleHeight && g.ball.y > g.Paddle2 {
		g.ball.vx *= -1
	}
}

func (g *GameState) wallCollision() {
	if g.ball.y < 0 || g.ball.y+BallHeight > Height {
		g.ball.vy *= -1
	}
	if g.ball.x < 0 || g.ball.x > Width {
		g.Reset()
	}
}

func (g *GameState) movePaddle1Up() {
	g.Paddle1 -= PaddleVelocity

	if g.Paddle1 < 0 {
		g.Paddle1 = 0

	}
}

func (g *GameState) movePaddle2Up() {
	g.Paddle2 -= PaddleVelocity
	if g.Paddle2 < 0 {
		g.Paddle2 = 0

	}
}

func (g *GameState) movePaddle1Down() {
	g.Paddle1 += PaddleVelocity
	if g.Paddle1+PaddleHeight > Height {
		g.Paddle1 = Height - PaddleHeight

	}
}
func (g *GameState) movePaddle2Down() {
	g.Paddle2 += PaddleVelocity
	if g.Paddle2+PaddleHeight > Height {
		g.Paddle2 = Height - PaddleHeight

	}
}
func (g *GameState) MovePaddle1Up() {
	g.Paddle1 -= PaddleVelocity

	if g.Paddle1 < 0 {
		g.Paddle1 = 0

	}
}
func (g *GameState) MovePaddle1Down() {
	g.Paddle1 += PaddleVelocity
	if g.Paddle1+PaddleHeight > Height {
		g.Paddle1 = Height - PaddleHeight

	}
}
func (g *GameState) MovePaddle2Up() {
	g.Paddle2 -= PaddleVelocity

	if g.Paddle2 < 0 {
		g.Paddle2 = 0

	}
}
func (g *GameState) MovePaddle2Down() {
	g.Paddle2 += PaddleVelocity
	if g.Paddle2+PaddleHeight > Height {
		g.Paddle2 = Height - PaddleHeight

	}
}
