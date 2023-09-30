package user

import "fmt"

type User struct {
	name string
}

func (u *User) SendMessage(message string) {
	fmt.Printf("%s received message: %s\n", u.name, message)
}

func NewUser(name string) *User {
	return &User{
		name: name,
	}
}
