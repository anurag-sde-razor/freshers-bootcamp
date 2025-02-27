package main

import "fmt"

func main(){
	arr:=[]int{1,2,3,4,5}
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}
	fmt.Printf("\n")
	fmt.Println("length:",len(arr))

	fmt.Printf("\n")
	arr=append(arr,98,99,97,96)
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}
	fmt.Printf("\n")
	fmt.Println("length:",len(arr))

	fmt.Printf("\n")
}