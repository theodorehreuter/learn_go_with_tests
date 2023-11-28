package clockface

import (
	"math"
	"time"
)

// A point represents a 2d cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// secondhand is the unit vector of the second hand of an analogue clock at time 't'
// represented as a Point
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} //scale
	p = Point{p.X, -p.Y}                                      // flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // translate
	return p
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func minutesInRadians(t time.Time) float64 {
	return math.Pi
}
