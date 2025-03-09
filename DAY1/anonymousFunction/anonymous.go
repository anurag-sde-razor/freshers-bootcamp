//inline function without name

package main

import "fmt"

func main(){
    result:=func(x,y int )int {
       return x+y
	} 
	fmt.Println(result(5,6))
}