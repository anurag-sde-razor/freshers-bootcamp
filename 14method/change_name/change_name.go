package main

import "fmt"

type Name struct {
	Name string
}

func (n *Name) chnageName(newName string) {
	n.Name = newName
}

func main() {
	name1 := Name{
		Name: "Anurag Aman",
	}
	fmt.Println("Name:", name1.Name)
	name1.chnageName("Abhishek Aman")
	fmt.Println("Name:", name1.Name)

}
