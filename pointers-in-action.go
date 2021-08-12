package main

import "fmt"

func main() {
	x:=0
	//foo(x)
	fmt.Println(x)
	//pointers
	z:=0
	fmt.Println("address of z", &z)
	fmt.Println("z before", z)
	foob(&z)
	fmt.Println("address of z after", &z)
	fmt.Println("z after", z)

}
/*
	Why pointers?
	Pointers are used when
1. you have a large chunk of data and you dont want to pass
	the data around your program in order to save memory and enhance performance,
	you just point at the address

2.  or when you need to change something at a certain location

	//Everything in Go is Pass by value.

*/
func foo(y *int)  {
	fmt.Println(y)
	//y=43
	fmt.Println(y)
	//
}
//with referencing
func foob(y *int)  {
	//fmt.Println(y)
	//y=43
	fmt.Println(y)
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y=43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}