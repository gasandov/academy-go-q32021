package routes

import (
	"github.com/gasandov/academy-go-q32021/controllers"

	"github.com/labstack/echo/v4"
)

func CreateEchoRoutes(e *echo.Echo) *echo.Echo {
	healthController := controllers.NewHealthController()
	pokemonController := controllers.NewPokemonController()
	consumeController := controllers.NewConsumerController()

	e.GET("/health-check", healthController.GetHealthCheck)

	e.GET("/pokemons", pokemonController.GetPokemons)
	e.GET("/pokemons/:id", pokemonController.GetPokemonById)

	e.GET("/consume", consumeController.ConsumeAPI)

	return e
}
