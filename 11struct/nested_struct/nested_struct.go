package main

import "fmt"

type employes struct{
     employinfo PersonalInfo 
	 employcontact Contact
	 employAddress personalAddres
}
type PersonalInfo struct{
	firstName  string
	lastName   string
	marriage   bool

}
type Contact struct{
    phone   string
	email   string
}

type personalAddres struct{
   TempAddress  string
   PermAddress string

}



func main(){

    employ1:=employes{
		employinfo:PersonalInfo{
			firstName: "Anurag",
			lastName: "Aman",
			marriage: false,
		},
		employcontact: Contact{
			phone : "+9189266971",
			email : "anurag.nitm@gmail.com",
		},
		employAddress: personalAddres{
			TempAddress: "jeevan bima nagar",
			PermAddress: "Chatarpur New Delhi",
		},
	}
   

   fmt.Println("employ personal info:",employ1.employinfo)
   fmt.Println("employ contact:",employ1.employcontact)
   fmt.Println("Address:",employ1.employAddress)


}