package basicDemo

var abc = "G"

var ab string
func main() {
	n()
	m()
	n()

	print("\n")

	ab = "G"
	print(ab)
	f1()
}

func n() { print(abc) }
func m() {
	a := "O"
	print(a)
}

func f1() {
	ab := "O"
	print(ab)
	f2()
}
func f2() {
	print(ab)
}