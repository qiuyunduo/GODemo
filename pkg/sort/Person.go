package sort

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

type Persons []Person

func (ps Persons) Len() int {
	return len(ps)
}

func (ps Persons) Less(a, b int) bool {
	return ps[a].lastName > ps[b].lastName
}

func (ps Persons) Swap(a, b int) {
	ps[a], ps[b] = ps[b], ps[a]
}

func PersonsTest() {
	persons := Persons{Person{"qiu", "yunduo"}, Person{"qiu", "yunhu"}, Person{"qiu", "yunjiao"}, Person{"qiu", "cheng"}}
	fmt.Println(persons)
	Sort(persons)
	fmt.Println(persons)
}
