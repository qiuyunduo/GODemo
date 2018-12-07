package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Rope string
func main() {
	 for i := 0; i < 10; i++ {
	 	a := rand.Int()
		 fmt.Printf("%d / ", a)
	 }
	fmt.Println()
	for i := 0; i < 5; i++ {
		a := rand.Intn(8)
		fmt.Printf("%d / ", a)
	}
	fmt.Println()

	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}

	fmt.Println()
	var name Rope = "qiuyunduo"
	fmt.Printf("%v", name)

	var ch0 byte = '\377'  //'\x41'
	fmt.Printf("%v",ch0)

	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3) // UTF-8 code point


	println(`This is a raw string \n`)
}
