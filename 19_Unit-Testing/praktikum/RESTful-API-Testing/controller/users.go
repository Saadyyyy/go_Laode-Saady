package controller

import (
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/config"
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/controller/request"
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/helper"
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/model"
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/repositorie"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []model.User

	if dbError := config.DB.Find(&users); dbError.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dbError.Error.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	users := model.User{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"massage": "Ivalid ID",
		})
	}

	if result := config.DB.First(&users, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "Get user by id Suscess",
		"User":    users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	var user model.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID",
		})
	}

	if result := config.DB.First(&user, id); result != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "User not found",
			})
		}
	}

	// Hapus pengguna dari database
	if err := config.DB.Unscoped().Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete user",
			"Data":    user,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
		"data":    user,
	})

}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	var users model.User
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"massage": "Invalid Input Id",
		})
	}

	//Mendapatkan data berdasarkan ID nya
	result := config.DB.First(&users, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "User not found",
			})
		}
	}

	// Persiapkan data yang ingin diupdate
	var updatedUser model.User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
		})
	}

	// Gunakan Updates tanpa memeriksa kolom kosong
	if err := config.DB.Model(&users).Updates(updatedUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update user",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "Update data suscessfully",
		"User":    users,
	})
}

//Login User
func LoginUserController(c echo.Context) error {
	var userRespon request.LooginRequest
	err := c.Bind(&userRespon)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponErr("Eror"))
	}

	data, token, err := repositorie.CheckLogin(userRespon.Email, userRespon.Password)

	UserResponse := map[string]interface{}{
		"token":   token,
		"user_id": data.ID,
		"name":    data.Name,
	}

	return c.JSON(http.StatusOK, helper.ResponSuccess("Succes Login", UserResponse))
}
