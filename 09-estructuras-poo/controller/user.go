package controller

import "time"

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

var users = make([]User, 0)

func GetUsers() []User {
	return users
}

func GetUser(id int) User {
	user := users[id]
	return user
}

func CreateUser(name string) {
	users = append(users, User{
		Id:        len(users) + 1,
		Name:      name,
		CreatedAt: time.Now(),
		Status:    true,
	})
}
