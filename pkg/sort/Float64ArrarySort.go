package sort

import (
	"fmt"
	"math/rand"
	"strings"
)

type Flocat64Arrary []float64

func (floats Flocat64Arrary) Len() int {
	return len(floats)
}

func (floats Flocat64Arrary) Less(a, b int) bool {
	return floats[a] > floats[b]
}

func (floats Flocat64Arrary) Swap(a, b int) {
	floats[a], floats[b] = floats[b], floats[a]
}

func NewFloat64Arrary() Flocat64Arrary {
	return make([]float64, 25)
}

func Fill() Flocat64Arrary {
	fs := make([]float64, 10)
	for i := 0; i < 10; i++ {
		fs[i] = (rand.Float64() * 100)
	}

	return fs
}

func (floats Flocat64Arrary) list() string {
	str := "{"

	for i := 0; i < floats.Len(); i++ {
		if floats[i] == 0 {
			continue
		}

		//str += strconv.FormatFloat(floats[i],'f',2,64)
		str += fmt.Sprintf("%4.1f , ", floats[i])
	}

	trim := strings.Trim(str, " , ")
	str = trim
	str += "}"
	return str
}

func (floats Flocat64Arrary) String() string {
	return floats.list()
}

func Flocat64Test() {
	arrary := NewFloat64Arrary()
	fmt.Println(arrary)
	fill := Fill()
	fmt.Println(fill)
	fmt.Println("The fill isSort :", IsSort(fill))

	Sort(fill)

	fmt.Println(fill)
	fmt.Println("The fill isSort :", IsSort(fill))
}
