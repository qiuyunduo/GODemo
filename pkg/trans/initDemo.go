package trans

import "fmt"

var Pi float64 = 0.00

type Person struct {
	name string
	age int
	sex string
}

func init()  {
	Pi = 3.14159256
}

func Test() int {
	return 5
}

func CreatePerson()  {
	person := Person{"qiuyunduo", 22, "man"}
	PrintPerson(&person)
}

func PrintPerson(person *Person) {
	fmt.Println((*person).name)
	fmt.Println((*person).age)
	fmt.Println((*person).sex)
}