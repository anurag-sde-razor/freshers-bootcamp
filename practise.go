package main
import "fmt"

func main(){
	var arr[3] int
	for index, _ :=range arr{
		fmt.Println("index:",index)
	}
}
