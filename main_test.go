package main

import (
	"testing"
)

func TestMyCalculator(t *testing.T) {
	if addition(2,3,1) != 6 {
		//fmt.Println("Expected 6, returned")
		t.Error("Expected 2+3+1 = 6")
	}

	if subtract(2,2) != 0 {
		t.Error("Expected 2-2 = 0")
	}

	if multiply(2,3,4)!=24 {
		t.Error("Expected 2*3*6 = 24")
	}

	if divide(2,2) != 1 {
		t.Error("Expected 2/2 = 1")
	}
}

func TestAddition(t *testing.T) {
	var additionStructs = [] struct {
		input [] float64
		expected float64
	}{
		{[]float64{1,2,3,8}, 14},
		{[]float64{3,5,6},14},
		{[]float64{2,3,4,3,3.5,7},22.5},
	}
for _, test:= range additionStructs {
	if output:=addition(test.input...); output != test.expected {
		t.Error("Test failed", test.input, test.expected, output)
	}
}

//func TestModify () {
//
//	}


}