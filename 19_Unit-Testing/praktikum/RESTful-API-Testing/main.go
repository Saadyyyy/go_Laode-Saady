package main

import (
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/config"
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/route"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	e := echo.New()
	route.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
