package rect

type Rectangular struct {
	length int
	width  int
}

type otherTest int

type Rect Rectangular

func (r *Rect) InitRect(length, width int) {
	r.length = length
	r.width = width
	MyTypeTest()
}

func (r Rect) InitRect1(length, width int) {
	r.length = length
	r.width = width
}

func (r *Rect) calArea() (area int) {
	area = r.length * r.width
	return
}
func (r *otherTest) calArea() (area int) {
	area = 12
	var i otherTest = 1
	*r = i
	return
}

func MyTypeTest() {
	test := new(otherTest)
	test.calArea()
	println(*test)
}

func (r *Rect) calPerimeter() (perimeter int) {
	perimeter = (r.length + r.width) * 2
	return
}

func (r Rect) Area() (area int) {
	area = r.calArea()
	return
}

func (r Rect) Perimeter() (perimeter int) {
	perimeter = r.calPerimeter()
	return
}
