package main

import "fmt"

type user struct {
	id       int
	username int // username bertipe int yang dimana kita tau kalau username bertipe string agar bisa di isi dengan value nama harusnya username bertipe string
	password int
}

// kita membuat lagi struct
type userservice struct {
	t []user
}

func (u userservice) getallusers() []user {
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}

func main() {
	u := user{
		id:       20,
		username: 334,
		password: 37382,
	}

	fmt.Println(u)
}
