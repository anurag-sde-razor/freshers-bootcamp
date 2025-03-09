package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func swap(a,b string)(string,string){
     return b,a
}

func main() {
	fmt.Println("random Number:", rand.Intn(10))
	fmt.Println("Square Root", math.Sqrt(20))
	fmt.Println("pi:", math.Pi)
	fmt.Println("Add", add(2, 7))
	fmt.Println("sub:", sub(9, 12))
	// fmt.Println("swap:",swap ("abhishek" ,"aman"))
	a, b := 5, 8
	c := add(a, b)
	fmt.Println("add", c)

	word1,word2:=swap("aman","abhishek")
	fmt.Println("After Swap",word1,word2)
}
