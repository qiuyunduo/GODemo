package basicDemo

import (
	"fmt"
	"runtime"
	"../../../pkg/trans"
)

func main() {


	i := trans.Test()
	fmt.Println(i,trans.Pi)
	fmt.Println("Hello, World!")
	fmt.Printf("%s", runtime.Version())
	print("sd")

	println("从前有座灵剑山")

}

func init()  {
	fmt.Println("Hello world!!!")
}
