package basicDemo

import "fmt"

func main() {
	var a,b,c int
	a,b,c = 1,2,2
	a,b = b,a
	fmt.Print("hello:",a,b,c)
	
}
