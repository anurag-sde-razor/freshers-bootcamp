package main

import "fmt"


type person struct{
	firstName string
	lastName string
	Age   int
}

func main(){
	var person1 person

	person1.firstName="Anurag"
	person1.lastName="Aman"
	person1.Age=22
	


fmt.Println("FirstName:",person1.firstName)
fmt.Println("LastNmae:",person1.lastName)
fmt.Println("Age:",person1.Age)

}