package main

import "fmt"

func main() {

	// pointer detail assigning

	var num int = 3

	var ptr *int //ye ptr ek address store krega kis tarah ki karega jis block me int tarah ki value hogi

	ptr = &num

	fmt.Println("value of num:", num)
	fmt.Println("value inside ptr:", ptr)
	fmt.Println("ptr continaing some address then value inside that address block:", *ptr)

	//pointer direct assigning and declaring

	num1 := 4
	ptr1 := &num1

	fmt.Println("value of ptr1:", *ptr1)

	//by default in pointer it conatin nil value

	var ptr2 *int
	if ptr2 == nil {
		fmt.Println("ptr2 is not assigned yet:")
	}
	num2:=55
	ptr2=&num2
	if ptr2 !=nil {
		fmt.Println("ptr2 is now assigned:",*ptr2)
	}


}
