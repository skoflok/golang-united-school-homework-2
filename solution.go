package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SideCounter int

const Pi float64 = math.Pi
const SidesTriangle SideCounter = 3
const SidesSquare SideCounter = 4
const SidesCircle SideCounter = 0

var availableSides = []SideCounter{SidesTriangle, SidesSquare, SidesCircle}

func CalcSquare(sideLen float64, sidesNum SideCounter) float64 {
	var square float64 = 0
	switch sidesNum {
	case 0:
		square = calcCircle(sideLen)
	case 3:
		square = caclTriangle(sideLen)
	case 4:
		square = calcSquare(sideLen)
	default:
		square = 0
	}
	return square
}

func caclTriangle(sideLen float64) float64 {
	var coef float64 = math.Pow(3, 0.5) / 4
	return coef * math.Pow(sideLen, 2)
}

func calcSquare(sideLen float64) float64 {
	return math.Pow(sideLen, 2)
}

func calcCircle(sideLen float64) float64 {
	return Pi * math.Pow(sideLen, 2)
}
