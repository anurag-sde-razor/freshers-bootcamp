package main

import "fmt"

func ByReference(a *int) {
	*a = *a - 2
}

func main() {
	num := 10
	ByReference(&num)
	fmt.Println("changed Number:", num)

}
