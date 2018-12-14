package interf

import (
	"fmt"
	"math/rand"
)

type any interface{}

type inter1 interface {
	sayHi()
}

type inter2 interface {
	talking()
}

type inter3 interface {
	inter1
	inter2
	stoping()
}

type Stru1 struct {
	index int
}

func (str *Stru1) sayHi() {
	fmt.Println("HI ........")
}

func (str *Stru1) talking() {
	fmt.Println("： How are you?")
	fmt.Println("： I am fine, Thanks")
}

func (str *Stru1) stoping() {
	fmt.Println("GoodBye.......")
}

func StrInterTest() {
	var i3 inter3
	i3 = &Stru1{12}
	i3.sayHi()
	i3.talking()
	i3.stoping()

	fmt.Println(i3.(*Stru1))

	var i1 inter1
	i1 = new(Stru1)
	i1.sayHi()

	fmt.Println(i1.(*Stru1))
	i1 = i3
	i1.sayHi()
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}

func TypeSwitch(any2 any) {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type\n", v)
		case int:
			fmt.Printf("any %v is an int type\n", v)
		case float64:
			fmt.Printf("any %v is a float32 type\n", v)
		case string:
			fmt.Printf("any %v is a string type\n", v)
		case Stru1:
			fmt.Printf("any %v is a special Stru1!\n", v)
		default:
			fmt.Println("unknown type!\n")
		}
	}
	testFunc(any2)
}

func TypeTest() {
	classifier(123, 123.123, true, nil, "Hello", Circle{})

	for i := 0; i < 10; i++ {
		f := rand.Float64()
		fmt.Println(f, " ")
	}

}

func AnyTest() {
	anies := []any{4, 5.2, "HelloWorld", true, Stru1{1}}

	for _, val := range anies {
		switch val.(type) {
		case int:
			fmt.Println("The is int")
		case string:
			fmt.Println("The is string")
		case bool:
			fmt.Println("The is bool")
		case float32:
			fmt.Println("the is float")
		case Stru1:
			fmt.Println("The is Stru1")
		default:
			fmt.Println("The is default")
		}
	}

	for _, val := range anies {
		TypeSwitch(val)
	}
}
