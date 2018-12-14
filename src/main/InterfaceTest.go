package main

import (
	"../../pkg/interf"
	"../../pkg/min"
	"../../pkg/point"
	"../../pkg/sort"
)

func main() {
	interf.ShaperTest()

	interf.TestReader()

	interf.TestSimplier()

	interf.ShapeTest()

	interf.StrInterTest()

	interf.TypeTest()

	interf.RSimpleTest()

	sort.SortTest()

	sort.PersonsTest()

	sort.Flocat64Test()

	point.PonintTest()

	interf.AnyTest()

	min.MinTest()
}
