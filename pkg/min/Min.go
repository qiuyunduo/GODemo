package min

import "fmt"

type Miner interface {
	Len() int
	At(index int) interface{}
	Less(a, b int) bool
	Swap(a, b int)
}

func Min(miner Miner) interface{} {
	min := miner.At(0)

	for i := 1; i < miner.Len(); i++ {
		if miner.Less(i, i-1) {
			miner.Swap(i, i-1)
		} else {
			min = miner.At(i)
		}
	}

	return min
}

type IntArrary []int

func (ints IntArrary) Len() int {
	return len(ints)
}

func (ints IntArrary) At(index int) interface{} {
	return ints[index]
}

func (ints IntArrary) Less(a, b int) bool {
	return ints[a] > ints[b]

}

func (ints IntArrary) Swap(a, b int) {
	ints[a], ints[b] = ints[b], ints[a]
}

func MinTest() {
	arraries := IntArrary{456, 234, 123123, 1231, 231, 23123, 123, 12312, 3123}
	fmt.Println(arraries)
	min := Min(arraries)
	fmt.Println("The arraries min element is ", min)
}
