package main

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

