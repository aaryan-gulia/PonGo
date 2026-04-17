package pong

import (
	"math"
)

const (
	Width                   float64 = 100
	Height                  float64 = 100
	PaddleHeight            float64 = 16
	PaddleWidth             float64 = 2
	PaddleVelocity          float64 = 1.33
	BallHeight              float64 = 2
	BallWidth               float64 = 2
	BallVelocityBase        float64 = 0.83
	BallVelocityMultiplier1 float64 = 1.5
	BallVelocityMultiplier2 float64 = 1.25
)

type GameEvent int

const (
	W GameEvent = iota
	S
	P
	Q
	N
)

type GameState struct {
	Paddle1  float64
	Paddle2  float64
	Points1  uint
	Points2  uint
	Ball     Ball
	hitCount uint
}

func (g *GameState) PollState() {
	if !g.paddleCollision() {
		g.wallCollision()
	}
	g.moveBall()
}

func (g *GameState) Reset() {
	g.Ball = initBall()
	g.Paddle1 = Height/2 - PaddleHeight/2
	g.Paddle2 = Height/2 - PaddleHeight/2
	g.hitCount = 0
}

func (g *GameState) moveBall() {
	g.Ball.step()
}

func (g *GameState) wallCollision() {
	if g.Ball.Y < 0 || g.Ball.Y+BallHeight > Height {
		g.Ball.invertY()
	}
	if g.Ball.X < 0 {
		g.Points2++
		g.Ball.reset()
		g.hitCount = 0
	}
	if g.Ball.X > Width {
		g.Points1++
		g.Ball.reset()
		g.hitCount = 0
	}
}

func (g *GameState) paddleCollision() bool {
	if g.Ball.X < PaddleWidth+BallWidth && g.Ball.Y < g.Paddle1+PaddleHeight && g.Ball.Y > g.Paddle1 {
		vec, ydir := computeCollisionBounce(g.Paddle1, g.Ball.Y)
		g.Ball.v.ydir = ydir
		g.Ball.v.xdir = positive
		g.Ball.v.VecComponents = vec
		g.hitCount++
		if g.hitCount == 4 {
			g.Ball.applyMultiplier(BallVelocityMultiplier1)
		}
		if g.hitCount == 12 {
			g.Ball.applyMultiplier(BallVelocityMultiplier2)
		}
		return true
	}
	if g.Ball.X > Width-PaddleWidth-BallWidth && g.Ball.Y < g.Paddle2+PaddleHeight && g.Ball.Y > g.Paddle2 {
		vec, ydir := computeCollisionBounce(g.Paddle2, g.Ball.Y)
		g.Ball.v.ydir = ydir
		g.Ball.v.xdir = negative
		g.Ball.v.VecComponents = vec
		g.hitCount++
		if g.hitCount == 4 {
			g.Ball.applyMultiplier(BallVelocityMultiplier1)
		}
		if g.hitCount == 12 {
			g.Ball.applyMultiplier(BallVelocityMultiplier2)
		}
		return true
	}
	return false
}

func computeCollisionBounce(paddle float64, y float64) (VecComponents, direction) {
	ydir := positive

	paddleZone := PaddleHeight / 8
	dif := paddle + PaddleHeight/2 - y
	if dif > 0 {
		ydir = negative
	}

	switch dif = math.Abs(dif); {
	case dif < paddleZone/2:
		return v1(), ydir
	case dif >= paddleZone/2 && dif < 2*paddleZone:
		return v2(), ydir
	case dif >= 2*paddleZone && dif < 3*paddleZone:
		return v3(), ydir
	case dif >= 3*paddleZone:
		return v4(), ydir
	}
	return v1(), ydir
}

//
///
////
/////
/////
////
///
//

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
