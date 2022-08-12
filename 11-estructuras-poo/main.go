package main

import (
	"fmt"
	"time"

	"github.com/Gilgammesh/go-practices/09-estructuras-poo/controller"
)

// Inheritance struct user
type UserInheritance struct {
	controller.User
}

func main() {
	// Create users
	controller.CreateUser("Carlos")
	controller.CreateUser("Marianella")
	controller.CreateUser("Jorge")
	controller.CreateUser("Pepe")
	controller.CreateUser("Martin")

	// Get user
	user := controller.GetUser(3)
	fmt.Printf("User %d: %s \n", user.Id, user.Name)

	// Get list all users
	users := controller.GetUsers()
	fmt.Println(users)

	// Inheritance user
	userInheritance := new(UserInheritance)
	userInheritance.Id = 1
	userInheritance.Name = "Max"
	userInheritance.CreatedAt = time.Now()
	userInheritance.Status = true
	fmt.Println(userInheritance)
}
