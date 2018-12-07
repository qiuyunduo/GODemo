package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.IntSize)
	i, e := strconv.Atoi("12")
	itoa := strconv.Itoa(i)
	fmt.Print(i,"    ",e,"   ",itoa,"\n")
	atoi, i2 := strconv.Atoi("123")
	fmt.Print(atoi,"   ",i2,"\n")
	float := strconv.FormatFloat(12.23231, 'b', 2, 64)
	fmt.Print(float)
}
