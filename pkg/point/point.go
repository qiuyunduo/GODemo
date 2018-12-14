package point

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Point3 struct {
	X, Y, Z float64
}

type Polar struct {
	R, T float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
func (p *Point) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

func (p *Point3) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z))
}

func (p Polar) Abs() float64 { return p.R }

//func Abs(p *Point) float64 {
//	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
//}
//
//func Scale(p *Point, s float64) (q Point) {
//	q.X = p.X * s
//	q.Y = p.Y * s
//	return
//}

func PonintTest() {
	p1 := new(Point)
	p1.X = 3
	p1.Y = 4
	m = p1 // p1 is type *Point, has method Abs()
	fmt.Printf("The length of the vector p1 is: %f\n", m.Abs())

	p2 := &Point{4, 5}
	m = p2
	fmt.Printf("The length of the vector p2 is: %f\n", m.Abs())

	p1.Scale(5)
	m = p1
	fmt.Printf("The length of the vector p1 after scaling is: %f\n", m.Abs())
	fmt.Printf("Point p1 after scaling has the following coordinates: X %f - Y %f\n", p1.X, p1.Y)

	mag := m.Abs()
	m = &Point3{3, 4, 5}
	mag += m.Abs()
	m = Polar{2.0, math.Pi / 2}
	mag += m.Abs()
	fmt.Printf("The float64 mag is now: %f\n", mag)
}
