package main
import "fmt"

func main(){
	arr:=[]int{1,2,3,4,5,6}
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}
}