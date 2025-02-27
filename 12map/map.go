package main

import "fmt"


func main(){
	map1:=make(map[string]int)

	map1["Anurag"]=89
	map1["Aman"]=100
	map1["rahul"]=79
	map1["harsh"]=58

	fmt.Println("Anurag Marks:",map1["Anurag"])

	// delete key value pair in map

	delete(map1,"Anurag")
	fmt.Println("Anurag Marks:",map1["Anurag"])
}