package main

import "fmt"

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	IsClose     bool
	Rating      float64
}

func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,

) User {
	return User{}
}

func (u *User) ChangeName(userName string) {
	if userName != "" {
		u.Name = userName
	}
}
func (u *User) ChangeAge(age int) {
	if age > 0 && age < 150 {
		u.Age = age
	}
}
func (u *User) ChangePhoneNumber(phoneNumber string) {
	if phoneNumber != "" {
		u.PhoneNumber = phoneNumber
	}
}

func main() {

	user := User{
		// Age:         10,
		Name:        "Jon",
		PhoneNumber: "+380923233",
		Rating:      0,
		IsClose:     false,
	}
	fmt.Println(user.Age)
	user.Age = 12
	fmt.Println(user.Age)

	fmt.Println(user.Rating, "Rating")
	fmt.Println(user.Age)
}
