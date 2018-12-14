package trans

import (
	"../../pkg/mystruct"
	"fmt"
	"reflect"
)

import (
	"../../pkg/file"
	"../../pkg/point"
	"../../pkg/rect"
	"strings"
	"time"
)

type T struct {
	i1  int
	f1  float64
	str string
}

type UserName struct {
	firstName string
	lastName  string
}

type Node struct {
	pre   *Node
	suf   *Node
	data  *UserName
	index int
}

type Address struct {
	country string "the user country"
	area    string "the user area"
}

type Vcard struct {
	name     string
	address  *Address
	birthday time.Time
	imageUrl string
}

type Annonymous struct {
	string
	int
	bool
	AnnonymousSon
}

type AnnonymousSon struct {
	float64
	float32
}

type node Node

func init() {

}

func upUserName(p *UserName) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func TestUserName() {
	// 1-struct as a value type:
	var pers1 UserName
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upUserName(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)
	// 2—struct as a pointer:
	pers2 := new(UserName)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward"
	// 这是合法的
	upUserName(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)
	// 3—struct as a literal:
	pers3 := &UserName{"Chris", "Woodward"}
	upUserName(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}

func StructTest() {
	t := new(T)
	t.i1 = 2
	t.f1 = 3.2212
	t.str = "Tony"

	fmt.Printf("The object t is : %v \n", t)

	InitStructT(t)

	fmt.Printf("The has be init, now object t is : %v \n", t)
}

func InitStructT(t *T) {
	//t = &T{12,123.12312,"Unknown"}
	//
	//t = &T{i1:23,f1:12.12,str:"Tom"}

	t.str = "Unknown"

	fmt.Printf("The has object t is : %v \n", t)
}

func TestNode() {
	mynode1 := &Node{nil, nil, nil, 0}
	mynode2 := &node{nil, nil, nil, 1}
	mynode3 := Node(*mynode2)
	mynode2.pre = mynode1
	mynode2.suf = &mynode3

	fmt.Println(mynode1, "  ", mynode2, "   ", &mynode3)
}

func VcardTest() {
	address := &Address{"China", "深圳市南山区东方科技大厦2115"}
	refTag(*address, 0)
	refTag(*address, 1)
	vcard := new(Vcard)
	vcard.name = "邱运铎"
	vcard.address = address
	birthday, _ := time.Parse("2006-01-02 15:04:05", "1998-03-15 08:37:18")
	fmt.Println(birthday)
	vcard.birthday = birthday
	vcard.imageUrl = "http://qiuyunduo.io/images/001.jpg"
	fmt.Println(vcard, "其中生日转化为string格式：", vcard.birthday.Format("2006-01-02 15:04:05"))
}

func RectTest() {
	myrect := new(rect.Rect)
	myrect.InitRect(12, 12)
	fmt.Println("The myrect area is ", myrect.Area())
	fmt.Println("The myrect Perimeter is ", myrect.Perimeter())
	myrect.InitRect1(1, 1)
	fmt.Println("The myrect area is ", myrect.Area())
	fmt.Println("The myrect Perimeter is ", myrect.Perimeter())
}

func PointTest() {
	p1 := new(point.Point)
	p1.X = 3
	p1.Y = 4
	fmt.Printf("The length of the vector p1 is: %f\n", point.Abs(p1))

	p2 := &point.Point{4, 5}
	fmt.Printf("The length of the vector p2 is: %f\n", point.Abs(p2))

	q := point.Scale(p1, 5)
	fmt.Printf("The length of the vector q is: %f\n", point.Abs(&q))
	fmt.Printf("Point p1 scaled by 5 has the following coordinates: X %f - Y %f \n", q.X, q.Y)
}

func FileTest() {
	file2 := file.NewFile(1, "001.txt")
	fmt.Println(file2)
	matrix := file.NewMatrix()
	fmt.Println(matrix)

}

func refTag(address Address, ix int) {
	addressType := reflect.TypeOf(address)
	field := addressType.Field(ix)
	fmt.Println(field.Tag)
}

func AnnonyMousTest() {
	an := new(Annonymous)
	an.string = "qiuyunduo"
	an.int = 22
	an.bool = true
	an.float32 = 32
	an.float64 = 64
	fmt.Println(an)

	annonymous := Annonymous{"123", 123, true, AnnonymousSon{23.23, 12.123}}
	fmt.Println(annonymous)
}

func EmployeeTest() {
	em := mystruct.NewEmployee("qiuyunduo", "java", 10000)
	fmt.Println("the before giveRaise", em)
	em.GiveRaise(1)
	fmt.Println("the after giveRaise", em)
}
