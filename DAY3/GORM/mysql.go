package main

import (
	// "github.com/go-sql-driver/mysql"
	//"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user struct {
	id        int      `gorm:"primaryKey";"autoIncrement"`
	FirstName string
	LastName  string
	Email     string
}

func main() {
	dsn := "root:pusHp@512@tcp(127.0.0.1:3306)/myDB"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // gorm.open() is a method to connect to the database
	if err != nil {
		panic("Connection Failed")
	}
	//fmt.Println(db)
	db.AutoMigrate(&user{})
	u := []user{
	{
		FirstName: "Rahul",
		LastName:  "Kumar",
		Email:     "rahul.kumar@gmail.com",
	},
	{
       FirstName: "Anurag",
	   LastName:"Aman",
	   Email:"anurag.aman@gmail.com",

	},
	{
		FirstName:"shivam",
		LastName:"kumar",
		Email:"shivam.kumar@gmail.com",
	},
	{
        FirstName:"sanjana",
		LastName: "Puri",
		Email:"sanjana.puri@gmail.com",
	},
}
	//db.Create(&u)
	//db.Updates(&u)
    
	// v:=user{
	// 	Email:"shivam.kumar@gmail.com",
	// }

    db.Where("email=?","sanjana.puri@gmail.com").Delete(&u)





}

