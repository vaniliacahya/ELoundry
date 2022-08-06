package main

import (
	"RESTAPILoundry/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := mysql.InitDB()
	mysql.MigrateData(db)
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())   //WAJIB!!
	e.Use(middleware.Logger()) //WAJIB!!
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	fmt.Println("E-loundry Running on....")
	e.Logger.Fatal(e.Start(":8000"))

}
