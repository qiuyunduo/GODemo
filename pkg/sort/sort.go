package sort

import "fmt"

type Sorter interface {
	Len() int
	Less(a, b int) bool
	Swap(a, b int)
}

type ListInt []int

func (listInt ListInt) Len() int {
	return len(listInt)
}

func (listInt ListInt) Less(a, b int) bool {
	return listInt[a] > listInt[b]
}

func (listInt ListInt) Swap(a, b int) {
	listInt[a], listInt[b] = listInt[b], listInt[a]
}

func Sort(data Sorter) {
	for i := 1; i < data.Len(); i++ {
		for j := 0; j < data.Len()-i; j++ {
			if data.Less(j, j+1) {
				data.Swap(j, j+1)
			}
		}
	}
}

func IsSort(data Sorter) bool {
	for i := 0; i < data.Len()-1; i++ {
		if data.Less(i, i+1) {
			return false

		}
	}
	return true
}

func SortTest() {
	ints := ListInt{1, 2, 32, 43, 54, 56, 6, 7, 2, 1, 3, 31, 3, 312}
	fmt.Println(ints)
	fmt.Println("The arrary isSort :", IsSort(ints))
	Sort(ints)
	fmt.Println(ints)
	fmt.Println("The arrary isSort :", IsSort(ints))
}
