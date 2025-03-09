package main

import (
	"fmt"
	"time"
)

func hello(){
	fmt.Println("Hello from hello function")
	time.Sleep(2*time.Second)
	fmt.Println("After 2 second Delay")
	fmt.Println("hello Anurag")
}

func demo(){
	fmt.Println("Hello from demo function")
}

func main() {
	fmt.Println("Goroutine_Started")
	go hello()
	time.Sleep(2000*time.Millisecond)
	//fmt.Println(hello("Anurag"))
	demo()
}
