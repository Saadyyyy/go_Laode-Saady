package main

type user struct {
	id       int
	username int // username bertipe int yang dimana kita tau kalau username bertipe string agar bisa di isi dengan value nama
	password int
}

// Untuk penamaan Struct nya lebih baik menggunakan Camel case
// userService
type userservice struct {
	t []user
}

// Penamaan pada method ini perlu di perbaiki karna tidak menetapkan clean code
// disini menurut saya meode yang cocok di gunakan adalah Camel case
func (u userservice) getallusers() []user { // getAllUsers
	return u.t
}

// Penamaan pada method ini perlu di perbaiki karna tidak menetapkan clean code
// disini menurut saya meode yang cocok di gunakan adalah Camel case
func (u userservice) getuserbyid(id int) user { // getUserById
	// untuk penamaan r kurang tepat karna susah membacanya lebih baik di ganti dengan `value`
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
