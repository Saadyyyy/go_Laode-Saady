package main

import (
	configs "saady/app/config"
	db "saady/app/database"
	mig "saady/app/migration"
	"saady/app/router"

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

func sum(n []int) []int {
	result := []int{}
	for i := len(n); i > 0; i-- {
		result = append(result, i+10)
	}
	return result
}
