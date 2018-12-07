package main

import (
	"fmt"
	"math"
	"runtime"
	"../../pkg/trans"
)

var prompt = "Enter a digit,e.g.3 " + "or %s to quit"

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt," ctrl+z Enter")
	}else {
		fmt.Println(runtime.GOOS)
		prompt = fmt.Sprintf(prompt,"ctrl+D")
	}
}

func main() {
	abs := math.Abs(-23)
	fmt.Println(abs)

	if trans.Test() > 7 {
		fmt.Println("Test() return > 7")
	} else {
		fmt.Println("Test() return < 7")
	}

	if val := 1; val > 0 {
		fmt.Println("This is if")
		fmt.Println(prompt)
	} else if true {
		fmt.Println("This is else if")
		return
	} else {

		fmt.Println("This is else")
	}

	var cond int = 9
	if cond = 20; cond > 10 {
		fmt.Println("This is cond > 10", cond)
	} else {
		fmt.Println("This si cond < 10", cond)
	}

	switch {
	case cond < 10:
		fmt.Println("this cond < 10")
	case cond > 10:
		fmt.Println("this cond > 10")
	default:
		fmt.Println("进入了default中")
	}

	switch cond = 30; cond {
	case 1, 2, 3, 4, 5, 6:
		fmt.Println("this is of 1,2,3,4,5,6")
	case 19, 20, 21, 22, 23, 24:
		fmt.Println("this is of 19,20,21,22,23,24")
	default:
		fmt.Println("this is default")

	}

	switch cond1 := trans.Test(); {
	case cond1 > 0:
		fmt.Println("this is of 1,2,3,4,5,6")
	case cond1 < 0:
		fmt.Println("this is of 19,20,21,22,23,24")
	default:
		fmt.Println("this is default")

	}

	month := 13
	if result, message := trans.Season(month); message == "ok" {
		fmt.Printf("This month %v is correct Season is %v.",month,result)
	} else {
		fmt.Printf("This month %v is error: %v",month,result)
	}
}
