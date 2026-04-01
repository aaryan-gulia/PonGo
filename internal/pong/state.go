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
	Up
	Down
)

type Ball struct {
	x  float64
	y  float64
	vx float64
	vy float64
}

type GameState struct {
	paddle1 float64
	paddle2 float64
	points1 int
	points2 int
	ball    Ball
}

func (g *GameState) PollState() {
	g.paddleCollision()
	g.wallCollision()
	g.moveBall()

}

func (g *GameState) HandleEvent(e GameEvent) {
	switch e {
	case W:
		g.movePaddle1Up()
	case S:
		g.movePaddle1Down()
	case Up:
		g.movePaddle2Up()
	case Down:
		g.movePaddle2Down()
	}
}

func (g *GameState) Reset() {
	g.ball.x = Width / 2
	g.ball.y = Height / 2
	g.ball.vx = BallVelocityBase
	g.ball.vy = BallVelocityBase
}

func (g *GameState) moveBall() {
	g.ball.x += g.ball.vx
	g.ball.y += g.ball.vy
}

func (g *GameState) paddleCollision() {
	if g.ball.x < PaddleWidth && g.ball.y < g.paddle1+PaddleHeight && g.ball.y > g.paddle1 {
		g.ball.vx *= -1
	}
	if g.ball.x > Width-PaddleWidth && g.ball.y < g.paddle2+PaddleHeight && g.ball.y > g.paddle2 {
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
	g.paddle1 -= PaddleVelocity

	if g.paddle1 < 0 {
		g.paddle1 = 0

	}
}

func (g *GameState) movePaddle2Up() {
	g.paddle2 -= PaddleVelocity
	if g.paddle2 < 0 {
		g.paddle2 = 0

	}
}

func (g *GameState) movePaddle1Down() {
	g.paddle1 += PaddleVelocity
	if g.paddle1+PaddleHeight > Height {
		g.paddle1 = Height - PaddleHeight

	}
}
func (g *GameState) movePaddle2Down() {
	g.paddle2 += PaddleVelocity
	if g.paddle2+PaddleHeight > Height {
		g.paddle2 = Height - PaddleHeight

	}
}
