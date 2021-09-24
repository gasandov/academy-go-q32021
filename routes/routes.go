package routes

import (
	"github.com/labstack/echo/v4"

	"academy-go-q32021/controllers"
)

func CreateEchoRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/health-check", controllers.HealthCheck)

	e.GET("/pokemons", controllers.GetPokemons)
	e.GET("/pokemons/:id", controllers.GetPokemonById)

	return e
}
