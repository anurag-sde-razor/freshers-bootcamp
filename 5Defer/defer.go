package main

import "fmt"

func mul(x, y int) int{
	res:=x*y
	fmt.Println("result:",res)
	return 0
}

func main(){
   
    fmt.Println("Starting..........")
	mul(8,5)
    defer mul(9,1)
    fmt.Println("hey1")
	fmt.Println("hey2")
	defer mul(3,2)
	fmt.Println("hey3")
	fmt.Println("hey4")

}