package controller

import (
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/config"
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/helper"
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/model"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

//Get all Blog
func GetBlogsController(c echo.Context) error {
	var blog []model.Blog

	if err := config.DB.Preload("User").Find(&blog); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Data Not Found",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "Get all Blogs",
		"Blogs":   blog,
	})
}

// create new Buku
func CreateBlogController(c echo.Context) error {
	Blog := model.Blog{}
	c.Bind(&Blog)

	if err := config.DB.Save(&Blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new Blog",
	})
}

//Get Blog By Id
func GetByIdController(c echo.Context) error {
	Blog := model.Blog{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponErrId(""))
	}

	if err := config.DB.Preload("User").First(&Blog, id); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, helper.ResponErr("Data Not Found"))
		}
	}

	return c.JSON(http.StatusOK, helper.ResponSuccess("Succes get blog by id", Blog))
}

//Update Blog
func UpdateBlogController(c echo.Context) error {
	Blog := model.Blog{}
	BlogUpdate := model.Blog{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponErrId(""))
	}

	if err := config.DB.First(&Blog, id); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, helper.ResponErr("Not Found"))
		}
	}

	if err := c.Bind(&BlogUpdate); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponErr("Internal Eroor"))
	}

	if err := config.DB.Model(&Blog).Updates(BlogUpdate); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, helper.ResponErr("Cannnot Update"))
		}
	}

	return c.JSON(http.StatusOK, helper.ResponSuccess("Success Update data", "Success"))
}

//Delete Blog
func DeleteBlogController(c echo.Context) error {
	Blog := model.Blog{}
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponErrId(""))
	}

	if err := config.DB.First(&Blog, id); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, helper.ResponErr("Not Found"))
		}
	}
	if err := config.DB.Delete(&Blog); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, helper.ResponErr("Cannot Delete Data"))
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "Success",
	})
}
