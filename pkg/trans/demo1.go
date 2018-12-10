package trans

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

func init() {
	fmt.Println("Hello")
}

func PrintArrary(str []string) {
	for _, value := range str {
		fmt.Print(value, "  ")
	}
	fmt.Println()
}

func PrintInts(arr []int) {
	for _, value := range arr {
		fmt.Print(value, "  ")
	}
	fmt.Println()
}

func Booltest() {
	avar := 10
	b := avar == 10
	c := avar == 5
	var a, d int = 1, 2
	d = a + a
	fmt.Println(d)
	fmt.Println(b, c)
	fmt.Println(3.56e2)
	fmt.Println()
}

func Season(month int) (result string, message string) {

	result, message = "错误的月份", "err"

	switch month {
	case 1, 2, 3:
		result, message = "春", "ok"
	case 4, 5, 6:
		result, message = "夏", "ok"
	case 7, 8, 9:
		result, message = "秋", "ok"
	case 10, 11, 12:
		result, message = "冬", "ok"
	}

	return result, message
}

func PrintG(hight int) {
	for i := 0; i < hight; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println(i)
	}
}

func PrintG1(hight int) {
	repeat := strings.Repeat("G", hight)
	for i := 1; i <= hight; i++ {
		s := string([]byte(repeat)[0:i])
		fmt.Println(s, "\t", i)
	}
}

func PrintRect(length int, width int) {
	for i := 0; i < width; i++ {
		switch {
		case i == 0, i == width-1:
			fmt.Println(strings.Repeat("*", length))
		case i > 0, i < width:
			repeat := strings.Repeat(" ", length-2)
			fmt.Println("*" + repeat + "*")
		default:
			println("err")
		}
	}
}

func PrintNumber100() {
	i := 0
LOOP:
	if i++; i < 100 {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print("FizzBuzz", " ")
		case i%3 == 0:
			fmt.Print("Fizz", " ")
		case i%5 == 0:
			fmt.Print("Buzz", " ")
		default:
			fmt.Print(i, " ")
		}
		goto LOOP
	}
	fmt.Println()
}

func PrintNumber10() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%b ", i)
	}
	fmt.Println()
}

func Forrange() {
	message := "How,are,you,!,I,am,fine,."

	fields := strings.Split(message, ",")

	for field := range fields {
		fmt.Print(fields[field], "\t")
	}

}

func SkipLabel() {
	goto LABELS
	fmt.Println("sdsd")
	println("asdasd")
LABELS:
	fmt.Println("我一下就到这了")
}

func Operint(num1, num2 int) (int, int, int) {
	return num2 + num1, num2 * num1, num2 - num1
}
func Operint1(num1, num2 int) (sum int, product int, difference int) {
	sum = num2 + num1
	product = num2 * num1
	difference = num2 - num1
	return sum, product, difference
}

func MethodChange(a, b int, reply *int) {
	*reply = a + b
}

func CreateStrings() {
	strs := []string{"hello", "world", "how", "are", "you"}
	PringStrings(&strs)
}

func PringStrings(strs *[]string) {
	for str := range *strs {
		fmt.Print((*strs)[str], " ")
	}
	fmt.Println()
}

func getFabonacci(index int) (num int) {
	if index == 1 || index == 2 {
		return 1
	} else {
		return getFabonacci(index-1) + getFabonacci(index-2)
	}
}

func getFabonacci1(index int) (index1 int, num int) {
	if index == 1 || index == 2 {
		return index, 1
	} else {
		return index, getFabonacci(index-1) + getFabonacci(index-2)
	}
}

func PringFabonacci(length int) {
	for i := 1; i <= length; i++ {
		_, num := getFabonacci1(i)
		fmt.Print(num, "  ")
	}
	fmt.Println()
}

func PringFabonacci1(length int) {
	start := time.Now()

	f := Fabnacci()
	for i := 1; i <= length; i++ {
		fmt.Print(f(i), "  ")
	}
	fmt.Println()

	end := time.Now()
	detal := end.Sub(start)
	fmt.Printf("this is no memorization fibernacci time : %s\n", detal)
}

func PrintNumber10Back(num int) {
	if num == 1 {
		fmt.Print(num, " ")
		fmt.Println()
		return
	}

	fmt.Print(num, " ")
	num--

	PrintNumber10Back(num)

}

func GetSum(numbers int) {
	fun := func() int {
		sum := 0
		for i := 0; i < numbers; i++ {
			sum += i
		}
		return sum
	}

	sum := fun()
	fmt.Printf("the 1000000 sum is =%d", sum)
	fmt.Println()
}

func Fabnacci() func(index int) int {
	x1, x2 := 0, 1
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

func ConsumeTime() {
	start := time.Now()
	for i := 1; i < 1000; i++ {
		fmt.Print("h", ",")
	}

	end := time.Now()
	detal := end.Sub(start)
	fmt.Println()
	fmt.Printf("the fun take this amouf of  time : %s\n", detal)
}

func Fibonacci_memonization(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = Fibonacci_memonization(n-2) + Fibonacci_memonization(n-1)
	}

	fibs[n] = res
	return
}

func PringFabonacci2(length int) {
	start := time.Now()

	for i := 0; i < length; i++ {
		fmt.Print(Fibonacci_memonization(i), "  ")
	}
	fmt.Println()

	end := time.Now()
	detal := end.Sub(start)

	fmt.Printf("this is use memorization fibernacci time : %s\n", detal)
}

func TestArrry() {
	b := [5]int{1, 2, 3, 4, 5}

	for i := range b {
		fmt.Print(i, " ")
	}
	fmt.Println()

	var arr1 = new([5]int)

	var arr2 = arr1
	arr2[1] = 100

	var arr3 = b
	arr3[1] = 12312312
	for i := 0; i < len(arr1); i++ {
		fmt.Print(arr1[i], " ")
	}
	fmt.Println()
	for i := 0; i < len(arr2); i++ {
		fmt.Print(arr2[i], " ")
	}
	fmt.Println()
	for i := 0; i < len(arr3); i++ {
		fmt.Print(arr3[i], " ")
	}
	fmt.Println()
	for i := 0; i < len(b); i++ {
		fmt.Print(b[i], " ")
	}
	fmt.Println()
}

func ReturnFabicaaiSlice(length int) []uint64 {
	var FibArr = [50]uint64{}
	FibArr[0] = 1
	FibArr[1] = 1

	for i := 2; i < len(FibArr); i++ {
		FibArr[i] = FibArr[i-1] + FibArr[i-2]
	}

	return FibArr[0:length]
}

func PrintFabicaaiArray(length int) {
	var FibArr = [50]uint64{}
	FibArr[0] = 1
	FibArr[1] = 1

	for i := 2; i < len(FibArr); i++ {
		FibArr[i] = FibArr[i-1] + FibArr[i-2]
	}

	for i := 0; i < length; i++ {
		fmt.Print(FibArr[i], " ")
	}
	fmt.Println()
}

func SliceTest() {
	var x [6]int
	var slice1 = x[0:3]

	for i := 0; i < len(x); i++ {
		x[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("The slice at %d is %d \n", i, slice1[i])
	}

	fmt.Printf("The x length is %d\n", len(x))
	fmt.Printf("The slice1 length is %d\n", len(slice1))
	fmt.Printf("The slice1 capital is %d\n", cap(slice1))

	slice2 := slice1[0:4]

	fmt.Printf("The slice2 length is %d\n", len(slice2))
	fmt.Printf("The slice2 length is %d\n", cap(slice2))

	//slice2等价于slice4
	slice3 := make([]int, 50, 100)
	slice4 := new([100]int)[0:50]
	fmt.Print(slice3[2], slice4[1], "\n")

	//切片为引用类型．改变切片中数值，其响应数组数值也会发生变化
	a := [...]string{"a", "b", "c", "d"}
	b := a[0:3]
	b[1] = "z"

	for i, value := range a {
		fmt.Println("Array item", i, "is", value)
	}

	for i := range b {
		fmt.Println("Array item", i, "is", b[i])
	}
}

func ByteBufferTest() {
	var buffer bytes.Buffer
	//var r *bytes.Buffer = new(bytes.Buffer)
	str := "123123"
	bts := []byte(str)

	bys := [3]byte{'s', 'd', 'q'}
	buffer.Write(bys[:])
	buffer.Write(bts)

	fmt.Printf("The buff add bys and now buff toString become: %s\n", buffer.String())

}

func Multidim_arrary() {
	arrs := [2][4]string{{"asd", "a", "s", "d"}, {"wio", "w", "i", "o"}}

	for i := 0; i < len(arrs); i++ {
		for j := 0; j < len(arrs[i]); j++ {
			fmt.Print(arrs[i][j], "  ")
		}
		fmt.Println()
	}

	fmt.Print(len(arrs), "   ", len(arrs[0]), "\n")
}

func getSliceFabicaai(length int) {
	slice := ReturnFabicaaiSlice(length)
	for i := range slice {
		fmt.Print(slice[i], "  ")
	}

	fmt.Println()

	fmt.Println("the slice capacity is : ", cap(slice))
}

func SliceTest1() {
	slice1 := make([]int, 0, 10)

	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Print(slice1[i], "  ")
	}
	fmt.Println()
}

func CopyAndAppendSlice() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)

	if n > cap(data) {
		newslice := make([]byte, (n+1)*2)
		copy(newslice, slice)
		slice = newslice
	}

	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func ExtendSlice(slice []int, factor int) []int {
	m := len(slice)
	n := m * factor

	if n > cap(slice) {
		newslice := make([]int, n, (n+1)*2)
		copy(newslice, slice)
		slice = newslice
	}

	slice = slice[0:n]
	return slice
}

func OperationSlice() {
	slice := make([]byte, 10, 50)
	appendByte := AppendByte(slice, 'w', 'e', 'e', 'd')
	fmt.Println(appendByte)
	appendByte[1] = 's'
	fmt.Println(slice)

	slice1 := make([]int, 10, 50)
	fmt.Printf("The now slice1 length is: %d\n", len(slice1))
	extendSlice := ExtendSlice(slice1, 4)
	slice1 = extendSlice
	fmt.Printf("The new slice1 length is: %d\n", len(slice1))

	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 10}
	fn := func(n int) bool {
		if n%2 == 0 {
			return true
		}

		return false
	}

	even := FilterEven(slice2, fn)
	fmt.Println(even)
	fmt.Println(len(even), "  ", cap(even))

	str1 := []byte{'q', 'w', 'e', 'r'}
	str2 := []byte{'h', 'e', 'l', 'l', 'o'}
	stringSlice := InsertStringSlice(str1, str2, 2)
	fmt.Println(string(stringSlice[:]))

	removeStringSlice := RemoveStringSlice(stringSlice, 2, 6)
	fmt.Println(string(removeStringSlice[:]))
}

func FilterEven(slice []int, fn func(int) bool) []int {
	evenSlice := make([]int, 0, len(slice))

	for _, value := range slice {
		if fn(value) {
			slice1 := append(evenSlice, value)
			evenSlice = slice1
		}
	}

	return evenSlice

}

func InsertStringSlice(toslice []byte, fromSlice []byte, index int) []byte {
	newslice := make([]byte, len(toslice)+len(fromSlice))
	prefixSlice := toslice[0:index]
	suffixSlice := toslice[index:]
	copy(newslice[0:index], prefixSlice)
	copy(newslice[index:index+len(fromSlice)], fromSlice)
	copy(newslice[index+len(fromSlice):], suffixSlice)
	return newslice
}

func RemoveStringSlice(slice []byte, start int, end int) []byte {
	newslice := make([]byte, len(slice)-(end-start)-1)
	copy(newslice[:start], slice[:start])
	copy(newslice[start:], slice[end+1:])
	return newslice
}

//if a == b -> 0 , if a > b -> 1 , if a < b -> -1
func CompareByte(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1

		}
	}

	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}

	return 0
}

func ByteArrTest() {
	arr1 := []byte("helloworld")
	arr2 := []byte("helloworld")

	compareByte := CompareByte(arr1, arr2)
	fmt.Println("the result arr1 compare with arr2 is:", compareByte)

	arr3 := []string{"q", "e", "ds"}
	sort.Strings(arr3)
	fmt.Println(arr3)

	i := append(arr3[:1], arr3[2:]...)
	fmt.Println(i)

	i2 := append(arr1[:2], append([]byte{'x'}, arr1[2:]...)...)
	fmt.Println(string(i2))

	i3 := append(arr2[:5], append(arr1, arr2[6:]...)...)
	fmt.Println(string(i3))
}

func ReadFile() {
	compile := regexp.MustCompile("[0-9]+")
	file, e := ioutil.ReadFile("doc/001.txt")
	fmt.Println(e)
	fmt.Println(string(file))

	find := compile.FindAll(file, len(file))

	for _, bytes := range find {
		fmt.Println(string(bytes))
	}
}

func SplitString(str string, index int) (a, b string) {
	i := append(make([]byte, 0, index), str[:index]...)
	i2 := append(make([]byte, 0, len(str)-index), str[index:]...)
	return string(i), string(i2)
}

func ReveserString(str string) string {
	bys := []byte(str)
	bys1 := make([]byte, len(str))

	for i := 0; i < len(str); i++ {
		bys1[i] = bys[len(str)-1-i]
	}

	return string(bys1)
}

func ReveserString1(str string) string {
	bys := []byte(str)
	for i := 0; i < len(str)/2; i++ {
		bys[i], bys[len(str)-i-1] = bys[len(str)-i-1], bys[i]
	}

	return string(bys)
}

func ReveserString2(str string) string {
	bys := []int32(str)
	for i := 0; i < len(str)/2; i++ {
		bys[i], bys[len(str)-i-1] = bys[len(str)-i-1], bys[i]
	}

	return string(bys)
}

func OperationString() {
	str := "helloworld"
	a, b := SplitString(str, 5)
	fmt.Println(a, "   ", b)

	s := str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(s)

	fmt.Println(ReveserString(str))
	fmt.Println(ReveserString1(str))
	fmt.Println(ReveserString2(str))
	CompareAndCopy([]byte("loop"))
}

func CompareAndCopy(chs []byte) []byte {
	fmt.Println("当前字符串.....")
	chs1 := make([]byte, 0, len(chs))
	for i := 0; i < len(chs); i++ {
		fmt.Print(string(chs[i]), " ")
		if i >= 1 && chs[i] != chs[i-1] {
			chs2 := append(chs1, chs[i])
			chs1 = chs2
		}
	}
	fmt.Println()
	fmt.Println("与前一个字符不相同的字符集合为：", string(chs1))
	return chs1
}

func OperationIntArr() {
	nums := []int{5, 34, 1, 78, 89, 5, 6, 7, 23}
	SortByBubbling(nums)
	fmt.Println(nums)

	f := func(n int) int {
		return n * 10
	}

	result := mapFun(f, 5, 34, 1, 78, 89, 5, 6, 7, 23)

	fmt.Println(result)
}

func SortByBubbling(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			if nums[j+1] < nums[j] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func mapFun(f func(n int) int, data ...int) []int {
	res := make([]int, len(data))
	for index := range res {
		res[index] = f(data[index])
	}

	return res

}

func CreateMap() {
	person := map[string]string{"name": "qiuyunduo", "age": "18", "sex": "man"}

	fnMap := map[int]func() int{
		1: func() int {
			return 10
		},
		2: func() int {
			return 20
		},
	}

	map1 := make(map[int][]int)

	map1[1] = []int{1, 2, 3, 4, 5, 6}

	_, ok := map1[2]
	fmt.Printf("The key is presist: %v", ok)

	for key, value := range person {
		fmt.Printf("The person %s is %s\n", key, value)
	}

	for key, value := range fnMap {
		fmt.Printf("The person %v is %v\n", key, value)
	}

}

func OpertatonMap() {
	CreateMap()
}
