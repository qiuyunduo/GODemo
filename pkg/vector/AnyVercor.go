package vector

type Element interface{}

type Vector struct {
	elements []Element
}

func (v *Vector) At(index int) Element {
	return v.elements[index]
}

func (v *Vector) Set(index int, element Element) {
	v.elements[index] = element
}
