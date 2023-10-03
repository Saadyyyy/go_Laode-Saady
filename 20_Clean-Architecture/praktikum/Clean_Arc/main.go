package main

import (
	configs "Laode_Saady/20_Clean-Architecture/praktikum/Clean_Arc/app/config"
	db "Laode_Saady/20_Clean-Architecture/praktikum/Clean_Arc/app/database"
	mig "Laode_Saady/20_Clean-Architecture/praktikum/Clean_Arc/app/migration"
	"Laode_Saady/20_Clean-Architecture/praktikum/Clean_Arc/app/router"

	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := configs.InitConfig()
	db := db.InitDbMysql(cfg)
	mig.MigrateDB(db)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVERPORT)))
}
