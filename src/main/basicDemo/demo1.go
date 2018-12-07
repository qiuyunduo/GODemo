//这是一个注释
//
//这是一个注释
//这是一个注释
package basicDemo

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import (
	"fmt";
	"os"
	"unsafe"
)

const pi float32 = 3.14159
const length  = len("123123")
const Ln2  = 0.69283728372372837921239398929038212
const beef, two, c  = "eat", 2, "veg"
const (
	Monday, Tuesday = 1, 2
	Wednesday, Thursday = 3, 4
)
const harEight = (1 << 100) >> 97

const (
	a = iota
	b
	d
)

var (
	HOME = os.Getenv("HOME")
)
func main() {
	a := 1
	println(a)
	fmt.Println(HOME)
	var num float32
	num = Ln2
	println(num)
	println(Random())
	println(Random())
	fmt.Println(test())
	fmt.Println(d)
	fmt.Println(pi)
	fmt.Println(length)
	fmt.Println(Ln2)
	fmt.Println(harEight)
	var s string = "sds"
	s = "sdsds"
	Print(s)
}

func Random() int {
	return int(C.random())
}

func test() (_ int,t int) {
	var j int
	j = 1
	return 2,j
}

func Print(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}
func FunctionName() int  {
	var i int
	i = 12
	return i;
}
