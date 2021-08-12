package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(subtract(4,3))
	fmt.Println(multiply(2,4,5,5))
	fmt.Println(subtractAll([]float64{3,5,6,1}))
	fmt.Println(calculation([]float64{20,2,2}))
	fmt.Println(Fact(4))
	fmt.Println(singleNumber([] int {4,1,2,1,2}))
	fmt.Println(math.Pow(float64(2), float64(3)))
	fmt.Println(3^5)
	//fmt.Println(subtractAll(-1-1-1-1))
}

func addition(nums ...float64) float64 {
	sum := float64(0)
	//for each index of num(element) within the range of nums
	for _, num := range(nums) {
		sum += num
	}
	return sum
}

//initialize return value to zero
// loop through the argument and add each

func multiply(nums ...float64) float64  {
	multiplied := float64(1)
	for _, num := range nums {
		multiplied*=num
	}
	return multiplied
}

func multiplyValues(nums ...float64) float64  {
	multiplied := float64(1)

	for i:=0; i<len(nums); i++ {
		multiplied = multiplied * float64(i)
		multiplied*= float64(i)
	}
	return multiplied
}

func subtract(i, j float64) float64  {
	sum:= i-j
	return sum
}

func divide(i, j float64) float64  {
	value:= i/j
	return value
}

func factorial(n int) int  {
	if n <= 1  {
		return 1
	} else {
		return n*factorial(n-1)
	}
}

func Fact(n int) int {
	if n<1 {
		return 1
	}
	fact:=n
	for n>2{
		n--
		fact*=Fact(n-1)
	}
	return fact
}
//func subtractAll(n ...float64) float64 {
//	/*
//	variable args,
//	 */
//	var value float64
//	for i,_:= range n{
//		//while n <= array length i.e (len(n)-1)::::Base case
//		if i <= len(n)-1 {
//		//	fmt.Println("check n[i]: ", n[i])
//			return n[i]
//		} else {
//			value=n[i]-(subtractAll(n[i]-1))
//			return value
//		}
//	}
//	return value
//}

// Do XOR of all elements and return
//int res = ar[0];
//for (int i = 1; i < ar_size; i++)
//res = res ^ ar[i];
//
//return res;
//}
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		math.Pow(float64(result), float64(nums[i]))
	}
	return result
}
func singleNumbers(nums []int) int {
	nums = [] int {4,1,2,1,2}
	var value = nums[0]
	//count:=0
	for i,_:= range nums{
		math.Pow(float64(value), float64(nums[i]))
		value^=nums[i]
		//if nums[i]==1 {
		//	value=nums[i]
		//	count++
		//}
	}
	return value
}


func subtractAll(n[]float64 ) float64 {
	//operator:= ""
	//
	//if operator=="*" {
	//	return subtractAll(n[:len(n)-1]) * n[len(n)-1]
	//} else if operator=="+" {
	//	return subtractAll(n[:len(n)-1]) + n[len(n)-1]
	//} else if operator=="-" {
	//	return subtractAll(n[:len(n)-1]) - n[len(n)-1]
	//} else {
	//	return subtractAll(n[:len(n)-1]) / n[len(n)-1]
	//}
		if len(n) == 1 {
			return n[0]
		}
			return subtractAll(n[:len(n)-1]) - n[len(n)-1]
}

func calculation(n[]float64) float64 {
	operator:= ""
	//strconv.FormatFloat([]n,32, 64,32)

	if len(n) == 1 {
		return n[0]
	} else {
		if operator=="*" {
			return calculation(n[:len(n)-1]) * n[len(n)-1]
		} else if operator=="+" {
			return calculation(n[:len(n)-1]) + n[len(n)-1]
		} else if operator=="-" {
			return calculation(n[:len(n)-1]) - n[len(n)-1]
		} else {
			return calculation(n[:len(n)-1]) / n[len(n)-1]
		}
	}
}

//func divideAll(n ...float64) float64  {
//	value:=float64(0)
//	for i,_:= range n{
//		if i<= len(n)-1{
//			return n[i]
//		} else {
//		//	value =	n[i] / (divideAll(n[i]+1))
//			value =	n / divideAll(n-1))
//			return value
//		}
//	}
//	return value
//}


//func calculate(s ...string) [] float64  {
//	eval:= make([]float64, 6, 10)
//	if strings.Contains(s,"*") {
//		strings.Split(s, "*")
//	} else if strings.Contains(s,"+") {
//		strings.Split(s, "+")
//	} else if strings.Contains(s,"-") {
//		strings.Split(s, "-")
//	} else {
//		strings.Split(s, "/")
//	}
//	len(s)
//	return eval
//}
//overall method
//func calculator(args ...string) (float64, error)   {
//
//	a, b, c, d  := "+", "-", "*", "/"
//	switch  {
//	case a:
//		fmt.Println(args, )
//
//
//	}
//
//
//	return strconv.ParseFloat(args, 64)
//}
/*
Write a program that takes in a variable number of strings, the character of each string should
be separated with the same type of operator e.g “2*4*5*5”, the operator should include;

· addition,

· subtraction,

· Multiplication and

· Division

The program must have its own function

It must take in a variable number of strings

It must return a slice of type float64 that contains the result of each operation

The length of the returned slice must be equal to the length of the argument passed in

Each operator should have their own functions that takes in a slice of type float64

They should have a return type of float64

· For example, calculate( “2*3*4*5” , “25/5/2”, “2+3+4+5.9+6+7.8”, “20.54-7.65-2.897”)
should return [120, 2.5, 28.7, 9.993]

· Divide [25,5,2] should return 2.5

· Each method must be tested, with a minimum of 80% test coverage.
 */

