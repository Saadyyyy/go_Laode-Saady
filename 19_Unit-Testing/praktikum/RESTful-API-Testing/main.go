package main

import (
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/config"
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/route"
)

func main() {
	e := route.New()
	config.Init()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
