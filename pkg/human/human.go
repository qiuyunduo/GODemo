package human

import (
	"fmt"
	"strconv"
)

type Men interface {
	SayHi()
	Singing()
}

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

type TwoInt struct {
	num1 int
	num2 int
}

func (man *Human) SetPhone(phone string) {
	man.phone = phone
}

func (man *Human) GetPhone() string {
	return man.phone
}

func (man *Human) Sleep() {
	fmt.Println("I am so tired ", man.name, " I will sleeping")
}

func (man *Human) SayHi() {
	fmt.Println("Hello I am ", man.name, " Hi!!!")
}

func (man *Human) Singing() {
	fmt.Println("I age is ", man.age, " phone is ", man.phone)
}

func (e *Employee) SayHi() {
	fmt.Println("I am working at ", e.company, " call me ", e.phone)
}

func (e *Employee) String() string {
	return "{" + e.name + "," + strconv.Itoa(e.age) + "," +
		e.phone + "," + e.company
}

/**
不要在String方法中调用涉及String的方法，例如fmt.Sprintf等，在调用String方法是会调用fmt.Sprintf，反过来fmt.Sprintf又会调用String，
形成无限递归调用,很快就会形成栈溢出
*/

func (tt TwoInt) String() string {
	return "(" + strconv.Itoa(tt.num1) + "/" + strconv.Itoa(tt.num2) + ")"
	//return fmt.Sprintf("%v",tt)  //error
}
func HumanTest() {
	milk := Student{Human{"milk", 22, "1232133211"}, "ecjtu"}
	tom := Employee{Human{"tom", 32, "1848233211"}, "cahubgke"}
	fmt.Println(&tom)

	milk.Sleep()
	tom.Sleep()

	var men Men
	fmt.Println("This is student milk")
	men = &milk
	men.SayHi()
	men.Singing()

	fmt.Println("This is Employee tom")
	men = &tom
	men.SayHi()
	men.Singing()

	fmt.Println("Tom phone is ", tom.GetPhone())
	tom.SetPhone("19029820902")
	fmt.Println("Now Tome phone is ", tom.GetPhone())

	twoInt := TwoInt{12, 23}
	tw := new(TwoInt)
	tw.num2 = 29
	tw.num1 = 22
	fmt.Printf("%v\n", twoInt)
	fmt.Println(tw)
}
