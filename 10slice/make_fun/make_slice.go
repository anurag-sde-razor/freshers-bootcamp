package main

import "fmt"

func main() {
	demo := make([]int, 3, 7)
	fmt.Println("demo slice", demo, len(demo), cap(demo))
	demo[0]=1
	demo[1]=2
	demo[2]=3
	demo=append(demo,23,45,67,8,98)
	// demo[0]=1
	// demo[0]=1
	fmt.Println("demo slice:", demo, "slice:",len(demo), "capacity:",cap(demo))


}
