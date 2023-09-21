package main

import (
	"Laode_Saady/17_ORM-Code-Structure-MVC/praktikum/config"
	"Laode_Saady/17_ORM-Code-Structure-MVC/praktikum/route"
)

func main() {
	e := route.New()
	config.Init()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
