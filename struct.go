package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
	first string
	ltk bool
}
func main() {
	//p1:= person{first: "Toluwase", last: "Thomas"}
	//p2:= person{
	//	last: "Alao",
	//}


	sa1:= secretAgent{
		person: person{
			first: "ab",
			last: "abc",
			age: 34,
		},
		first: "teem",
		ltk: true,
	}
	fmt.Println(sa1)

	fmt.Println(sa1.first, sa1.person.first)

}

//A