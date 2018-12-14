package main

import (
	"../../pkg/trans"
	"fmt"
	"log"
	"path"
	"runtime"
)

import . "sort"

var where func() = func() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
}

func main() {
	hight := 10
	trans.PrintG(hight)
	trans.PrintG1(hight)
	trans.PrintRect(100, 10)
	trans.PrintNumber100()
	where()
	trans.PrintNumber10()
	trans.Forrange()
	trans.SkipLabel()
	fmt.Println(trans.Operint(5, 3))
	fmt.Println(trans.Operint1(5, 2))
	trans.CreateStrings()
	trans.CreatePerson()
	trans.PringFabonacci(20)
	trans.PrintNumber10Back(10)
	trans.GetSum(1e6)

	trans.PringFabonacci1(41)
	trans.PringFabonacci2(41)

	trans.ConsumeTime()

	trans.TestArrry()

	trans.PrintFabicaaiArray(2)

	trans.SliceTest()

	trans.ByteBufferTest()

	trans.Multidim_arrary()

	trans.SliceTest1()

	trans.CopyAndAppendSlice()

	trans.OperationSlice()

	trans.ByteArrTest()

	trans.ReadFile()

	trans.OperationString()

	trans.OperationIntArr()

	trans.OpertatonMap()

	trans.RebootComputer()

	trans.OperationUnsaft()

	trans.OperationList()

	trans.RegexpTest()

	trans.LockTest()

	trans.BigIntTest()
	fmt.Println(path.Base("doc/001.txt"))

	Ints([]int{1, 2, 3, 4, 5, 6})
}
