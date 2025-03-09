package main

import "fmt"

func main() {
	message := make(chan string)
	message <- "Hello I am Anurag"
	fmt.Println(<-message)

}
