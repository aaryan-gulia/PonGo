package pong

import (
	"math/rand"
)

type direction int

type VecComponents struct {
	i float64
	j float64
}

func v1() VecComponents {
	return VecComponents{0.83, 0.00}
}
func v2() VecComponents {
	return VecComponents{0.83, 0.28}
}
func v3() VecComponents {
	return VecComponents{0.83, 0.55}
}
func v4() VecComponents {
	return VecComponents{0.83, 1.11}
}

type Vec struct {
	VecComponents
	xdir direction
	ydir direction
}

const (
	positive direction = 1
	negative direction = -1
)

type Ball struct {
	v      Vec
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
		v:      Vec{VecComponents: v2(), xdir: xdir, ydir: ydir},
		vscale: BallVelocityMultiplier0,
		X:      Width / 2,
		Y:      Height / 2,
	}
}

func (b *Ball) reset() {
	b.v.VecComponents = v2()
	ydir := positive
	if rand.Float64() < 0.5 {
		ydir = negative
	}
	b.v.ydir = ydir
	b.vscale = BallVelocityMultiplier0
	b.X = Width / 2
	b.Y = Height / 2
}

func (b *Ball) step() {
	b.X += float64(b.v.xdir) * b.v.i * b.vscale
	b.Y += float64(b.v.ydir) * b.v.j * b.vscale
}

func (b *Ball) applyMultiplier(m float64) {
	b.vscale = m
}

func (b *Ball) invertY() {
	b.v.ydir *= -1
}

func (b *Ball) invertX() {
	b.v.xdir *= -1
}
