package pong

const (
	Width                  float64 = 100
	Height                 float64 = 100
	PaddleHeight           float64 = 16
	PaddleWidth            float64 = 2
	PaddleVelocity         float64 = 3
	BallHeight             float64 = 2
	BallWidth              float64 = 2
	BallVelocityBase       float64 = 0.5
	BallVelocityMultiplier float64 = 1.5
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
	g.paddleCollision()
	g.wallCollision()
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

func (g *GameState) paddleCollision() {
	if g.Ball.X < PaddleWidth && g.Ball.Y < g.Paddle1+PaddleHeight && g.Ball.Y > g.Paddle1 {
		g.Ball.invertX()
		g.hitCount++
		if g.hitCount%4 == 0 && g.hitCount < 20 {
			g.Ball.applyMultiplier(BallVelocityMultiplier)
		}
	}
	if g.Ball.X > Width-PaddleWidth && g.Ball.Y < g.Paddle2+PaddleHeight && g.Ball.Y > g.Paddle2 {
		g.Ball.invertX()
		g.hitCount++
		if g.hitCount%4 == 0 && g.hitCount < 15 {
			g.Ball.applyMultiplier(BallVelocityMultiplier)
		}
	}
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
