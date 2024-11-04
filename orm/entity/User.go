package entity

import "fmt"

type User struct {
	Name string
	age  int
}

func NewUser(name string) User {
	return User{Name: name, age: 0}
}

func (u User) GetName() string {
	fmt.Println("getName-action！")
	return u.Name
}

func (u User) GetCustomName(name string) string {
	fmt.Println("GetCustomName！")
	u.Name = name
	return u.Name
}
