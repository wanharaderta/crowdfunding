package main

import (
	"crowdfunding/user"
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

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Test Save from service"
	userInput.Email = "contoh@gmail.com"
	userInput.Occupation = "Occupation"
	userInput.Password = "pass"

	userService.RegisterUser(userInput)
	//user := user.User{
	//	Name: "Test User Created",
	//}
	//userRepository.Save(user)
}
