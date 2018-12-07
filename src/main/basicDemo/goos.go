package basicDemo

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	a := 5
	var b int = a
	println(*&b)
	println(&a)
	goos := runtime.GOOS
	fmt.Printf("The operation system is: %v\n",goos)
	path := os.Getenv("PATH")
	sprintf := fmt.Sprintf("%s%v%d", "23", 23, 23)
	fmt.Println(sprintf)
	fmt.Printf("Path is %s\n",path)
	fmt.Printf("%s%v%d","23",23,23)

	fmt.Print("Hello:",23,23,23,23)

}
