package pong

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
	P
	Q
	N
)

type Ball struct {
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
	Ball    Ball
}

func (g *GameState) PollState() {
	g.paddleCollision()
	g.wallCollision()
	g.moveBall()

}

func (g *GameState) Reset() {
	g.Ball.X = Width / 2
	g.Ball.Y = Height / 2
	g.Ball.vx = BallVelocityBase
	g.Ball.vy = BallVelocityBase
}

func (g *GameState) moveBall() {
	g.Ball.X += g.Ball.vx
	g.Ball.Y += g.Ball.vy

}

func (g *GameState) paddleCollision() {
	if g.Ball.X < PaddleWidth && g.Ball.Y < g.Paddle1+PaddleHeight && g.Ball.Y > g.Paddle1 {
		g.Ball.vx *= -1
	}
	if g.Ball.X > Width-PaddleWidth && g.Ball.Y < g.Paddle2+PaddleHeight && g.Ball.Y > g.Paddle2 {
		g.Ball.vx *= -1
	}
}

func (g *GameState) wallCollision() {
	if g.Ball.Y < 0 || g.Ball.Y+BallHeight > Height {
		g.Ball.vy *= -1
	}
	if g.Ball.X < 0 || g.Ball.X > Width {
		g.Reset()
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
