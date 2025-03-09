package main
import "fmt"

func main(){
  // using range for slice
	arr:=[]int{1,2,3,4,5,6}

	for index,value:=range arr{
		fmt.Println("index:",index,"value:",value)
	}
	arr=append(arr,9,45,255,107)
	fmt.Printf("\n")

	for index,value:=range arr{
		fmt.Println("index:",index,"value:",value)
	}

	//using range for map
	fmt.Printf("\n")

	map1:=map[string]int{
		"Anurag":67,
		"rahul":89,
		"Mahraj":90,

	}
	for key,value :=range map1{
		fmt.Println("key:",key,"value:",value)
	}

}