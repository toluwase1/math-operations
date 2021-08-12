package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	&: Get the pointer: points to the address Ampersand Address
	*: Reference
	x:=7 | Value is 7, while Reference is the location where 7 is stored eg 0xc000018068
	Reference meaning where exactly is 7 stored in computer memory, not x now: this: 0xc000018068
	The value of a pointer is the address of another value in memory
	A pointer points to alocation in memory where value is stored

	A pointer to where an int is stored is a pointer to an int
	*/
	x:=7
	y:=&x
	fmt.Println(&y, "first")
	//&x is a pointer
	fmt.Println(&x)
	fmt.Println(x,y)
	//de-referencing 1: having access to the value stored in that address
	fmt.Println(*y)
	fmt.Println(*&x, "second")
	//de-referencing 2: manipulating the value stored in that address
	*y=8
	fmt.Println(x,y)
	//& gives you the address
	//* gives you the value stored in that address
	fmt.Println(*&x, "third")
	//s:=6
	//abc(&s)
	fmt.Println(abc)


	a:=42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	//when you have an address, the type of the address is a pointer to whatever value type is stored in the address
	//eg *int
	fmt.Printf("%T", &a)
	//var b int = &a



}

func abc(y *int, s string)  {
	*y= int(math.Pow(2, 2))
	fmt.Println("-------------")
	fmt.Println(*y, "*y")
	fmt.Println(y,"y")
}
