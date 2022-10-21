package main

import (
	"simple-app/database"
	"simple-app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()

	app := echo.New()

	routes.InitRoutes(app)

	app.Logger.Fatal(app.Start(":1323"))
}
