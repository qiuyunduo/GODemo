package main

import (
	"../../pkg/trans"
	"fmt"
	"log"
	"runtime"
)

var where func() = func() {
	_, file, line,_ := runtime.Caller(1)
	log.Printf("%s:%d",file,line)
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
	fmt.Println(trans.Operint(5,3))
	fmt.Println(trans.Operint1(5,2))
	trans.CreateStrings()
	trans.CreatePerson()
	trans.PringFabonacci(20)
	trans.PrintNumber10Back(10)
	trans.GetSum(1e6)

	f := trans.Fabnacci()
	for i := 1; i <= 20; i++ {
		fmt.Print(f(i),"  ")
	}
}

