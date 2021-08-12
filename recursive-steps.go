package main

//func main() {
//	fmt.Println(sum(5))
//	fmt.Println(multiplier(5))
//}

//1. Write recursive function that given an input n, program sums all non-negative integers up to n
//steps in solving a recursion problem
// Recursion involves solving a problem with a simpler version of the problem

// *STEP1* whats the simplest possible input for the function
// simplest input is  n=0, this wil be the base case for the problem.
//the base case is the only case where we provide explicit answer,
//all other solution to the problem will build on the base case

//STEP2: pLAY AROUND WITH EXAMPLES AND VISUALISE

//STEP3 RELATE HARDER EXAMPLES TO SMALLER EXAMPLES

// STEP4 GENERALIZE THE PATTERM

//WRITE THE CODE

func sum(n float64) float64 {
	if n==0 {
		return n
	} else {
		return n+sum(n-1)
	}
}
//if n==0, return 0; else if n=1, return 1+0;
//else if n=2, return 2+1+0; else if n==3, return 3+2+1+0

func multiplier(n float64) float64  {
	if n==1 {
		return n
	} else {
		return n*multiplier(n-1)
	}
}