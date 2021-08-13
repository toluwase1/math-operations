package main
var a = "G"
func main() {
	n()//G
	n()//G
	m()//0
	n()//G
	n()//G
}
func n() {
	print(a)
}
func m() {
	a = "O"
	print(a)
}