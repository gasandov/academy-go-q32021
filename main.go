package main

import (
	"log"

	"github.com/gasandov/academy-go-q32021/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e = routes.CreateEchoRoutes(e)

	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
