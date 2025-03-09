package main
import "fmt"

func main(){
	var arr[4] int
	for i:=0;i<len(arr);i++{
		var num int
		fmt.Scan(&num)
		arr[i]=num
	}
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}

	fmt.Printf("\n")
	//2nd way of decring as well as initialising

	arr1:=[3]int{1,2,3}
	fmt.Println(arr1[2])
}