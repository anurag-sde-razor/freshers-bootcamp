package main

import "fmt"

var sum int=0;

func Add(n int)int {
   for i:=1;i<=n;i++{
	sum=sum+i
   }
   fmt.Println("sum of n Number start from 1:",sum)
   return 0
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello World")
	}
   Add(10)
	

}
