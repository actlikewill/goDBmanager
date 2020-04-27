package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=postgres dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	dbase := db.DB()
	defer dbase.Close()

	err = dbase.Ping()
	if err != nil {
		panic(err.Error())
	}
	
	println("Connection to database established!!")
	db.DropTable(&User{})
	db.CreateTable(&User{})

	for _, user := range users {
		db.Create(&user)
	}

	u := User{Username: "tmacmillan"}
	db.Where(&u).First(&u)
	fmt.Println(u)

	u.LastName = "Beeblebrox"
	
	db.Save(&u)

	user := User{}
	db.Where(&u).First(&user)
	fmt.Println(user)

	db.Where(&User{Username: "adent"}).Delete(&User{})

	println("done")
}

// User represents a user object.
type User struct {
	ID uint
	Username string
	FirstName string
	LastName string

}

var users []User = []User{
	{Username: "adent", FirstName: "Arthur", LastName: "Dent"},
	{Username: "fprefect", FirstName: "Ford", LastName: "Prefect"},
	{Username: "tmacmillan", FirstName: "Tricia", LastName: "MacMillan"},
	{Username: "mrobot", FirstName: "Marvin", LastName: "Robot"},
}