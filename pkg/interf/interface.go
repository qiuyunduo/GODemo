package interf

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
)

//Go 语言中的接口都很简短,通常它们会包含 0 个、最多 3 个方法。

type Shaper interface {
	Area() float64
}

type Simpler interface {
	Get() interface{}
	Set(n interface{})
}

type RSimple struct {
	name string
}

type Simple struct {
	index int
}

func (s *Simple) Get() interface{} {
	return s.index
}

func (s *Simple) Set(n interface{}) {
	i := n.(int)
	s.index = i
}

func (rs *RSimple) Get() interface{} {
	return rs.name
}

func (rs *RSimple) Set(name interface{}) {
	i := name.(string)
	rs.name = i
}

type Shape struct {
}

type Square struct {
	side float64
}

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	Shape
	r float64 "圆的半径"
}

func (square *Square) Area() float64 {
	return math.Pow(square.side, 2)
}

func (rect *Rectangle) Area() float64 {
	return (rect.width + rect.length) * 2
}

func (shape Shape) Area() float64 {
	fmt.Println("The area is I don't know")
	return 0.12
}

func getArea(shaper Shaper) float64 {
	return shaper.Area()
}

func OperSimpler(simpler Simpler) {
	fmt.Println("The simpler index is ", simpler.Get())

	simpler.Set(12)
	fmt.Println("After update, the simpler index is ", simpler.Get())
}

func CheckType(simpler Simpler) {
	switch t := simpler.(type) {
	case *Simple:
		fmt.Println("The simpler is Simple and value :", t)
	case *RSimple:
		fmt.Println("The simpler is RSimple and value :", t)
	default:
		fmt.Println("The is default ", t)

	}
}

func ShaperTest() {
	square := new(Square)
	square.side = 5
	fmt.Println("The Square area is ", square.Area())
	fmt.Println("The Square area is ", getArea(square))

	rect := new(Rectangle)
	rect.length = 5
	rect.width = 3

	var shaper Shaper
	shaper = square

	shapers := make([]Shaper, 3, 5)
	shapers[0] = square
	shapers[1] = rect
	shapers[2] = shaper
	for _, val := range shapers {
		//fmt.Println("The shapers , and it is ",val.Area())
		fmt.Println("The shape area is ", getArea(val))
		if t, ok := val.(*Rectangle); ok {
			fmt.Println("..........")
			fmt.Println(t)
			fmt.Println("..........")
		}
	}

}

func TestReader() {
	var reader io.Reader
	reader = os.Stdin
	reader = bufio.NewReader(reader)
	reader = new(bytes.Buffer)
	file, _ := os.Open("doc/001.txt")
	reader = bufio.NewReader(file)
	i := make([]byte, 44, 4096)
	reader.Read(i)
	fmt.Println(string(i))

}

func TestSimplier() {
	var simpler Simpler
	simpler = &Simple{3}
	OperSimpler(simpler)
}

func ShapeTest() {
	shapers := [3]Shaper{&Square{2}, &Rectangle{4, 5}, &Circle{Shape{}, 2}}

	for _, val := range shapers {
		fmt.Println("The shaper are is ", val.Area())

		switch t := val.(type) {
		case *Square:
			fmt.Println("the is square", t, "  ", t.Area())
		case *Rectangle:
			fmt.Println("The is rectagle", t, "  ", t.Area())
		case *Circle:
			fmt.Println("The is Circle ", t, "  ", t.Area())

		default:
			fmt.Println("The is default ", t, "  ", t.Area())
		}
	}
}

func RSimpleTest() {
	simplers := []Simpler{new(Simple), new(RSimple)}
	CheckType(simplers[0])
	CheckType(simplers[1])
	simple := Simple{12}
	simple.Set(12)
	va, ok := simplers[0].(Simpler)
	fmt.Println(va, ok, simple)
}
