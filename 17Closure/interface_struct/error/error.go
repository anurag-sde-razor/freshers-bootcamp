package main

import "fmt"

func divide(x, y int) (int, error) {

	if y == 0 {
		return 0, fmt.Errorf("Cannot divide by z")
	}
	return x / y, nil

}

func main() {

	result, error := divide(10, 0)

	if error != nil {
		fmt.Println(error)
	} else {       // error is nil here
		fmt.Println(result)
	}

}
