package main

import (
	"crowdfunding/user"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connecting to database...")

	var users []user.User

	db.Find(&users)

	length := len(users)
	fmt.Println("Length of users", length)

	for _, user := range users {
		fmt.Println("User name", user.Name)
	}
}
