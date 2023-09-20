package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	GetId := User{}
	cek := false
	for _, v := range users {
		if v.Id == id {
			GetId = v
			cek = true
			break
		}
	}

	if cek == false {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "User not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by id",
		"user":     GetId,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	var deleteId *User
	cek := -1
	for i, v := range users {
		if v.Id == id {
			deleteId = &users[i]
			cek = i
			break
		}
	}

	if cek == -1 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"massage": "Data tidak ada",
		})
	}

	users = append(users[:cek], users[cek+1:]...)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "Succsess delete data from user",
		"User":    deleteId,
	})
}

// // update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	updatedUser := User{}
	c.Bind(&updatedUser)

	var updatedId *User

	cek := -1

	for i, v := range users {
		if v.Id == id {
			updatedId = &users[i]
			cek = i
			break
		}
	}
	if cek == -1 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"massage": "Data not found",
		})
	}
	updatedId.Name = updatedUser.Name
	updatedId.Email = updatedUser.Email
	updatedId.Password = updatedUser.Password

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "Succsess Updated data from user",
		"User":    updatedId,
	})

}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)

	//Need Solution Code
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))

}
