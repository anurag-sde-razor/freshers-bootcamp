package main

import "fmt"

var age int = 9

func switch1(x int) string {
	switch x {

	case 1:
		return "one"
	case 2:
		return "two"
	default:
		return "switch case Not Matched"
	}
}

func main() {
	if age > 18 {
		fmt.Println("Hey! You are Adult")
	} else {
		fmt.Println("Sorry! you are Minor")
	}
	result := switch1(2)

	fmt.Println(result)

}
