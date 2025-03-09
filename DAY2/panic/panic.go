package main

import "fmt"


func panic_demo(x,y int)int{
	if y==0{
		panic("Invalid Input!")
	}
	return x/y
}

func main(){
     fmt.Println("valid:",panic_demo(10,4))
	 fmt.Println("valid:",panic_demo(10,2))
	 fmt.Println("valid:",panic_demo(10,0))
	 fmt.Println("valid:",panic_demo(10,5))
	 fmt.Println("valid:",panic_demo(10,1))
	 
	 
}
