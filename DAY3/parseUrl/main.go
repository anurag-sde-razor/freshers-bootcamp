package main

import (
	"fmt"
	"net/url"
)

func main() {

	myURL := "https://jsonplaceholder.typicode.com/todos"
	fmt.Printf("type of Myurl: %T\n", myURL)
	result,err:=url.Parse(myURL)
	if err!=nil{
		fmt.Println("Error occured",err)
		return
	}

	 fmt.Printf("type of result: %T\n", result)

     fmt.Println("scheme",result.Scheme)
	 

	 fmt.Println("DominName:",result.Host)

}
