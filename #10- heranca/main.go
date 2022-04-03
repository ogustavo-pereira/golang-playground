package main

import "fmt"

type User struct {
	name   string
	online bool
}

func (u *User) Login() {
	u.online = true
}

func (u *User) Logout() {
	u.online = false
}

type Admin struct {
	User
	ranking int
}

func main() {
	var adm Admin

	adm.Login()
	fmt.Println(adm.online)

	adm.Logout()
	fmt.Println(adm.online)
}
