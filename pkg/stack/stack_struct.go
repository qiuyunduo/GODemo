package stack

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	index int
	data  [ARR_LIMIT]int
}

func (stack *Stack) Init() {
	stack.index = 0
}

func (stack *Stack) Push(val int) {
	if stack.index >= len(stack.data) {
		fmt.Println("The stack is full...")
		return
	}

	stack.data[stack.index] = val
	stack.index++

}

func (stack *Stack) Pop() int {
	if stack.index <= 0 {
		fmt.Println("The stack is empty")
		return 0
	}

	stack.index--
	return stack.data[stack.index]
}

func (stack *Stack) String() string {
	str := ""
	for i := 0; i < stack.index; i++ {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(stack.data[i]) + "] "
	}
	return strings.Trim(str, " ")
}
