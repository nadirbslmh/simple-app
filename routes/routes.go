package routes

import (
	"simple-app/controllers"

	"github.com/labstack/echo/v4"
)

var bookController controllers.BookController = controllers.BookController{}

func InitRoutes(e *echo.Echo) {
	bookRoute := e.Group("/api/v1/books")

	bookRoute.GET("", bookController.GetAll)
	bookRoute.GET("/:id", bookController.GetByID)
	bookRoute.POST("", bookController.Create)
	bookRoute.PUT("/:id", bookController.Update)
	bookRoute.DELETE("/:id", bookController.Delete)
}
