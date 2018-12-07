package trans

import (
	"fmt"
	"strings"
)

func init()  {
	fmt.Println("Hello")
}

func Booltest() {
	avar := 10
	b := avar == 10
	c := avar == 5
	var a,d int = 1,2
	d = a + a
	fmt.Println(d)
	fmt.Println(b,c)
	fmt.Println(3.56e2)
	fmt.Println()
}

func Season(month int) (result string,message string) {

	result,message = "错误的月份","err"

	switch month {
	    case 1,2,3:
	    	result,message = "春","ok"
	case 4,5,6:
		result,message = "夏","ok"
	case 7,8,9:
		result,message = "秋","ok"
	case 10,11,12:
		result,message = "冬","ok"
	}

	return result,message
}

func PrintG(hight int)  {
	for i := 0; i < hight; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println(i)
	}
}

func PrintG1(hight int)  {
	repeat := strings.Repeat("G", hight)
	for i := 1; i <= hight; i++ {
		s := string([]byte(repeat)[0:i])
		fmt.Println(s,"\t",i)
	}
}

func PrintRect(length int, width int)  {
	for i := 0; i < width; i++ {
		switch {
		case i == 0,i == width-1:
			fmt.Println(strings.Repeat("*", length))
		case i > 0,i < width:
			repeat := strings.Repeat(" ", length-2)
			fmt.Println("*"+repeat+"*")
		default:
			println("err")
		}
	}
}

func PrintNumber100()  {
	i := 0
	LOOP:
		if i++; i < 100 {
			switch {
			case i % 3 == 0 && i % 5 == 0:
				fmt.Print("FizzBuzz"," ")
			case i % 3 == 0:
				fmt.Print("Fizz"," ")
			case i % 5 == 0:
				fmt.Print("Buzz"," ")
			default:
				fmt.Print(i," ")
			}
			goto LOOP
		}
	fmt.Println()
}

func PrintNumber10()  {
	for i := 0; i < 10; i++ {
		fmt.Printf("%b ",i)
	}
	fmt.Println()
}

func Forrange()  {
	message := "How,are,you,!,I,am,fine,."

	fields := strings.Split(message, ",")

	for field := range fields {
		fmt.Print(fields[field],"\t")
	}

}

func SkipLabel()  {
	goto LABELS
	fmt.Println("sdsd")
	println("asdasd")
LABELS:
	fmt.Println("我一下就到这了")
}

func Operint(num1,num2 int) (int,int,int) {
	return num2+num1,num2*num1,num2-num1;
}
func Operint1(num1,num2 int) (sum int,product int, difference int) {
	sum = num2+num1
	product = num2*num1
	difference = num2-num1
	return sum,product,difference
}

func MethodChange(a,b int , reply *int)  {
	*reply = a + b
}

func CreateStrings()  {
	strs := []string{"hello","world","how","are","you"}
	PringStrings(&strs)
}

func PringStrings(strs *[]string)  {
	for str := range *strs {
		fmt.Print((*strs)[str]," ")
	}
	fmt.Println()
}

func getFabonacci(index int) (num int) {
	if index == 1 || index == 2 {
		return 1
	} else {
		return getFabonacci(index-1) + getFabonacci(index - 2)
	}
}

func getFabonacci1(index int) (index1 int,num int) {
	if index == 1 || index == 2 {
		return index,1
	} else {
		return index,getFabonacci(index-1) + getFabonacci(index - 2)
	}
}

func PringFabonacci(length int)  {
	for i := 1; i <= length; i++ {
		fmt.Print(getFabonacci(i),"  ")
	}
	fmt.Println()
}

func PrintNumber10Back(num int)  {
	if num == 1 {
		fmt.Print(num," ")
		fmt.Println()
		return
	}

	fmt.Print(num," ")
	num--

	PrintNumber10Back(num)

}

func GetSum(numbers int)  {
	fun := func() int{
		sum := 0
		for i := 0; i < numbers; i++ {
			sum += i
		}
		return sum
	}

	sum := fun()
	fmt.Printf("the 1000000 sum is =%d",sum)
	fmt.Println()
}

func Fabnacci()  func(index int) int{
	x1,x2 := 0,1
	sum := 0

	return func(index int) int {
		if index == 0 || index == 1 {
			return 1
		}
		sum = x1 + x2
		x1 = x2
		x2 = sum
		return sum
	}

}