package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("what is your name:")
    reader:=bufio.NewReader(os.Stdin)
	name,_:=reader.ReadString('\n')
	fmt.Println("name:",name )
}