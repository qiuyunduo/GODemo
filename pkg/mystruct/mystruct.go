/**
 * 类型和作用在它上面定义的方法必须在同一个包里定义,这就是为什么不能在 int、float 或类似这些的类型上定义方
 *法
 */

package mystruct

import (
	"fmt"
	"reflect"
	"strconv"
)

type employee struct {
	name       string
	occupation string
	salary     float64
}

type T struct {
	int
	float32
	string
}

type Day int

type TZ int

type IntVector []int

const (
	MO Day = iota
	TU
	WE
	TH
	FR
	SA
	SU
)

const (
	HOUR TZ = 60 * 60
	UTC  TZ = 0 * HOUR
	EST  TZ = -5 * HOUR
	CST  TZ = -6 * HOUR
)

var timezones = map[TZ]string{
	UTC: "University Greenwich time",
	EST: "Eastern Standard Time",
	CST: "Central Standard Time",
}

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (d Day) String() string {
	return dayName[d]
}

func (tz TZ) String() string {
	for tz1, zone := range timezones {
		if tz == tz1 {
			return zone
		}
	}

	return ""
}

func NewEmployee(name string, occuption string, salary float64) (em *employee) {
	em = new(employee)
	em.name, em.occupation, em.salary = name, occuption, salary
	return
}

func (e *employee) GiveRaise(factor float64) {
	e.salary += e.salary * factor
}

func (v IntVector) sum() (sum int) {
	sum = 0
	for _, value := range v {
		sum += value
	}

	return
}

func (list IntVector) Len() int {
	return len(list)
}

func (list *IntVector) Append(val int) {
	*list = append(*list, val)
}

func (t *T) String() string {
	return strconv.Itoa(t.int) + " / " + strconv.FormatFloat(float64(t.float32), 'f', 6, 32) + " / " + t.string
}

func MyStructTest() {
	fmt.Println("the sum of the IntVector is:", IntVector{2, 1, 2, 3, 54, 65, 7}.sum())
	t := &T{12, 23.23, "abc\\tdef"}
	fmt.Printf("%v\n", t)

	var th Day = 3
	fmt.Printf("The 3rd day is: %s\n", th)
	// If index > 6: panic: runtime error: index out of range
	// but use the enumerated type to work with valid values:
	var day = SU
	fmt.Println(day) // prints Sunday
	fmt.Println(0, MO, 1, TU)

	i := SU * 2
	fmt.Println(reflect.TypeOf(i))

	fmt.Println(EST)
	fmt.Println(0 * HOUR)
	fmt.Println(-6 * HOUR)
}

func ListTest() {
	var list IntVector
	list.Append(2)
	fmt.Println("the list ", list, " length is ", list.Len())

	vector := new(IntVector)
	vector.Append(8)
	fmt.Println("the vector ", vector, " length is ", vector.Len())
}
