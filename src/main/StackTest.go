package main

import (
	"../../pkg/stack"
	"fmt"
	"runtime"
)

func main() {
	stack1 := new(stack.Stack)
	stack1.Init()

	stack1.Pop()
	for i := 1; i < 10; i++ {
		stack1.Push(i)
	}

	fmt.Println(stack1)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}
