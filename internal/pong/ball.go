package pong

import "math"

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
