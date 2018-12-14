package main

import (
	"../../pkg/evenutil"
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Is the integer %d even? %v\n", i, evenutil.Even(i))
	}
}
