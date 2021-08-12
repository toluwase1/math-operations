package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(subtractVal([]float64{100,50, 10}))
	fmt.Println(addVal([]float64{100,50, 10}))
	fmt.Println(multiplyVal([]float64{100,50, 10}))
	fmt.Println(divideVal([]float64{100,50, 10}))
}

func subtractVal(n[]float64 ) float64 {
	if len(n) == 1 {
		return n[0]
	}
	return subtractVal(n[:len(n)-1]) - n[len(n)-1]
}

func multiplyVal(n[]float64 ) float64 {
	if len(n) == 1 {
		return n[0]
	}
	return multiplyVal(n[:len(n)-1]) * n[len(n)-1]
}

func divideVal(n[]float64 ) float64 {
	if len(n) == 1 {
		return n[0]
	}
	return divideVal(n[:len(n)-1]) / n[len(n)-1]
}

func addVal(n[]float64 ) float64 {
	if len(n) == 1 {
		return n[0]
	}
	return addVal(n[:len(n)-1]) + n[len(n)-1]
}

func calculateAll(s string) []float64 {
	//convert string to float64
	if strings.Contains(s, "*") {
		strings.Split(s,"*")
		strings.IndexRune(s,9)
		//floatValue := strconv.FormatFloat(3.1415, 'E', -1, 64)
		//addVal()
	}

	//perform operation (check if strings contain *,/,+,-)
	//return [] float64
return calculateAll("")
}

