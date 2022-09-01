package router

import (
	"test/controller"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func SetupRoute() *echo.Echo {
	e := echo.New()
	e.GET("/books/:id", controller.GetBook)        //grtting the specific books by the index
	e.POST("/books", controller.CreateBook)        //create the book
	e.DELETE("/delete/:id", controller.DeleteBook) //delete the book by index
	e.PUT("/updatebook/:id", controller.UpdateBook)
	return e
}
