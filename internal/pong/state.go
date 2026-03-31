package pong

const (
	Width                  float64 = 100
	Height                 float64 = 100
	PaddleHeight           float64 = 16
	PaddleWidth            float64 = 2
	BallHeight             float64 = 2
	BallWidth              float64 = 2
	BallVelocityBase       float64 = 1
	BallVelocityMultiplier float64 = 2
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
	if g.ball.x > 100-PaddleWidth && g.ball.y < g.paddle2+PaddleHeight && g.ball.y > g.paddle2 {
		g.ball.vx *= -1
	}
}

func (g *GameState) wallCollision() {
	if g.ball.y < 0 || g.ball.y > 100 {
		g.ball.vy *= -1
	}
	if g.ball.x < 0 || g.ball.x > 100 {
		g.Reset()
	}
}

func (g *GameState) paddleRebound() {
	//TODO
}
