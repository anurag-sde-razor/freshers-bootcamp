package main

import "fmt"

type person struct {
	name   string
	number string
	email  string
}

func main() {
	// 1st way of implementing Struct
	var person0 person
	person0.name="Abhishek Aman"
	person0.number="+91836007193"
	person0.email="abhishek.aman@gmail.com"

	fmt.Println("name",person0.name)
	fmt.Println("number",person0.number)
	fmt.Println("email",person0.email)
	fmt.Printf("\n")

	// 2nd way of implementing Struct
	person1 := person{
		name:   "Anurag Aman",
		number: "+918368865174",
		email:  "anurag.nitm@gmail.com",
	}
	fmt.Println("Name:",person1.name)
	fmt.Println("Number:",person1.number)
	fmt.Println("email:",person1.email)
	fmt.Printf("\n")

	//3rd way dynamically implementing struct using new Keyboard

	var person2 =new(person)// person2 store it as a pointer 
	person2.name="Harsh"
    person2.number="+918857294718"
	person2.email="harsh123@gmail.com"

	fmt.Println("person2--->",person2)

	fmt.Println("name:",person2.name)
	fmt.Println("number",person2.number)
	fmt.Println("email",person2.email)

}
