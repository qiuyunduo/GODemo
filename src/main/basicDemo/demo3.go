package basicDemo
var abcd = "G"
func main() {
	n1()
	m1()
	n1()
}
func n1() {
	print(abcd)
}
func m1() {
	abcd = "O"
	print(abcd)
}