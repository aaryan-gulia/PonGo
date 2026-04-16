package pong

import (
	"math"
	"math/rand"
)

type direction int

type unitVecComponents struct {
	i float64
	j float64
}

func makeUnitVecComponent(x float64, y float64) unitVecComponents {
	l2 := x*x + y*y

	return unitVecComponents{
		i: math.Sqrt(x * x / l2),
		j: math.Sqrt(y * y / l2),
	}

}

func v1() unitVecComponents {
	return makeUnitVecComponent(1, 0)
}
func v2() unitVecComponents {
	return makeUnitVecComponent(3, 1)
}
func v3() unitVecComponents {
	return makeUnitVecComponent(3, 2)
}
func v4() unitVecComponents {
	return makeUnitVecComponent(3, 4)
}

type unitVec struct {
	unitVecComponents
	xdir direction
	ydir direction
}

const (
	positive direction = 1
	negative           = -1
)

type Ball struct {
	v      unitVec
	vscale float64
	X      float64
	Y      float64
}

func initBall() Ball {
	xdir := positive
	if rand.Float64() < 0.5 {
		xdir = negative
	}
	ydir := positive
	if rand.Float64() < 0.5 {
		ydir = negative
	}
	return Ball{
		v:      unitVec{unitVecComponents: v2(), xdir: xdir, ydir: ydir},
		vscale: BallVelocityBase,
		X:      Width / 2,
		Y:      Height / 2,
	}
}

func (b *Ball) reset() {
	b.v.unitVecComponents = v2()
	ydir := positive
	if rand.Float64() < 0.5 {
		ydir = negative
	}
	b.v.ydir = ydir
	b.vscale = BallVelocityBase
	b.X = Width / 2
	b.Y = Height / 2
}

func (b *Ball) step() {
	b.X += float64(b.v.xdir) * b.v.i * b.vscale
	b.Y += float64(b.v.ydir) * b.v.j * b.vscale
}

func (b *Ball) applyMultiplier(m float64) {
	b.vscale *= m
}

func (b *Ball) invertY() {
	b.v.ydir *= -1
}

func (b *Ball) invertX() {
	b.v.xdir *= -1
}
