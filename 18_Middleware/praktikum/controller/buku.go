package controller

import (
	"Laode_Saady/18_Middleware/praktikum/config"
	"Laode_Saady/18_Middleware/praktikum/model"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetBukuController(c echo.Context) error {
	var Bukus []model.Buku

	if err := config.DB.Find(&Bukus).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all Bukus",
		"Bukus":   Bukus,
	})
}

// create new Buku
func CreateBukuController(c echo.Context) error {
	Buku := model.Buku{}
	c.Bind(&Buku)

	if err := config.DB.Save(&Buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new Buku",
		"Buku":    Buku,
	})
}

//Get buku by id
func GetBukuIdController(c echo.Context) error {
	buku := model.Buku{}
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"massage": "Invalid Id",
		})
	}
	if result := config.DB.First(&buku, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "Get Data Book By Id Success",
		"Buku":    buku,
	})

}

//Detele Book
func DeletedBukuController(c echo.Context) error {
	var buku model.Buku
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"massage": "Invalid Id",
		})
	}
	if result := config.DB.First(&buku, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Data Not Found",
			})
		}
	}

	if err := config.DB.Unscoped().Delete(&buku, id); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"massage": "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "Success Deleted Book",
		"Book":    buku,
	})
}

//Update Book
func UpdatedBukuController(c echo.Context) error {
	var buku model.Buku
	var bukuUpdate model.Buku
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"massage": "Invalid Id",
		})
	}

	if result := config.DB.First(&buku, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Data Not Found",
			})
		}
	}

	if err := c.Bind(&bukuUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid",
		})
	}

	if err := config.DB.Model(&buku).Updates(bukuUpdate).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"massage": "Cannot Update Data Book",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "Succes Updated Buku",
		"Buku":    buku,
	})
}
