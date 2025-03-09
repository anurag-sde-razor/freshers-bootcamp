// package main

// import "fmt"


// func recover1(x,y int)(int,int,int){
//     sum:=x+y
// 	mul:=x*y
// 	if y==0{
// 		panic("Invalid Input of Denominator!")
// 	}
// 	div:=x/y

// 	return sum,mul,div

// }

// func main(){

	

// 	 sum, mul, div := recover1(10, 2)
// 	 fmt.Println("Result:", sum, mul, div)
// 	 sum, mul, div = recover1(10, 3)
// 	 fmt.Println("Result:", sum, mul, div)
// 	 sum, mul, div = recover1(10, 0)
// 	 fmt.Println("Result:", sum, mul, div)
// 	 sum, mul, div = recover1(10, 4)
// 	 fmt.Println("Result:", sum, mul, div)

// 	 defer func(){
// 		if r:=recover();r!=nil{
// 			fmt.Println("Pnic Recovered:",r)
// 		}
// 	 }()

// }




package main

import "fmt"

package main

import "fmt"

// recover1 performs some basic operations and handles potential panics
func recover1(a, b int) (int, int, int) {
    // Simulate panic when b is 0
    if b == 0 {
        panic("Invalid Input of Denominator!")
    }
    return a + b, a * b, a / b
}

func main() {
    // Deferring the recovery function to catch panics in the main function
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Panic Recovered:", r)
        }
    }()

    sum, mul, div := recover1(10, 2)
    fmt.Println("Result:", sum, mul, div)

    sum, mul, div = recover1(10, 3)
    fmt.Println("Result:", sum, mul, div)

    // This will cause a panic
    sum, mul, div = recover1(10, 0)
    fmt.Println("Result:", sum, mul, div)

    // This will print, even after the panic
    sum, mul, div = recover1(10, 4)
    fmt.Println("Result:", sum, mul, div)
}
