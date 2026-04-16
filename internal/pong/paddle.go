package pong

type paddleZone uint

const (
	zoneOne paddleZone = iota
	zoneTwo
	zoneThree
	zoneFour
	zoneFive
	zoneNone
)

func getZone(y float64, p float64) paddleZone {
	switch {
	case y >= p && y < p+PaddleHeight/8:
		return zoneOne
	case y >= p+PaddleHeight/8 && y < p+3*PaddleHeight/8:
		return zoneTwo
	case y >= p+3*PaddleHeight/8 && y < p+5*PaddleHeight/8:
		return zoneThree
	case y >= p+5*PaddleHeight/8 && y < p+7*PaddleHeight/8:
		return zoneFour
	case y >= p+7*PaddleHeight/8 && y < p+PaddleHeight:
		return zoneFive
	default:
		return zoneNone
	}
}

//
///
///
////
////
///
///
//
