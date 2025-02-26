package main

import "fmt"

var a int
var b, c, d bool

var f, g int = 2, 6

var h float64=float64(10)/3

var i float32=float32(10)/3



func main() {
	v1:="Anurag Aman"
	rahul, ischild, price := "brother", true, 30000
	fmt.Println(a, b, c, d)
	fmt.Println(f, g)
	fmt.Println(rahul, ischild, price)
	fmt.Println("10/3",h)
	fmt.Println("10/3",i)
	fmt.Printf("v1 is of type %T\n", v1)
}
