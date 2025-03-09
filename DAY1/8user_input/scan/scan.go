package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("user-input...")
	var a, b int

	fmt.Println("Enter Number 1")
	fmt.Scan(&a)
	fmt.Println("Enter Number 2")
	fmt.Scan(&b)

	fmt.Println("Result:", add(a, b))
}
