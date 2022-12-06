package main

import (
	"fmt"
	"os"
	"simple-app/database"
	"simple-app/routes"

	"github.com/labstack/echo/v4"
)

const DEFAULT_PORT = "8080"

func main() {
	database.InitDB()

	app := echo.New()

	routes.InitRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = DEFAULT_PORT
	}

	appPort := fmt.Sprintf(":%s", port)

	app.Logger.Fatal(app.Start(appPort))
}
