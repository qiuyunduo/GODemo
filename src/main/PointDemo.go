package main

import "fmt"

func main() {
	s := 123
	var p *int = &s
	*p = 12
	fmt.Println(p,s,*p)
	
}
