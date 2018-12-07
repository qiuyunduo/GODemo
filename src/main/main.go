package main

import (
	"fmt"
	"../../pkg/trans"
)
func main() {
	trans.Booltest()
	fmt.Println("Sd")

	var num1 int16 = 43
	var num2 int32
	num2 = int32(num1)
	fmt.Println(num2)

	c := 5 + 10i
	c1 := 5 - 10i
	c2 := c * c1
	print(real(c),imag(c))
	println()
	print(real(c2),imag(c2))
	print(c2)
	fmt.Printf("%v%f", c, real(c))
	fmt.Printf("kesd"+"sdsds")
}
