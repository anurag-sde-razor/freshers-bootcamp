package main

import "fmt"

type people struct{
	name string
	age int
	contact string
}



func main(){

   var people1 people
   people1.name="Anurag Aman"
   people1.age=21 
   people1.contact= "anurag.nitm@gmail.com"


   people2:=people{

	  name:"Abhishek",
	  age:21,
	  contact:"abhishek.amam@gmail.com",

   }


   

   

  fmt.Println("name:",people1.name,"age:",people1.age,"contact:",people1.contact)
  fmt.Println("name:",people2.name,"age:",people2.age,"contact:",people2.contact)
  

}