package routes

import (
	"github.com/gasandov/academy-go-q32021/controllers"
	"github.com/gasandov/academy-go-q32021/repositories"
	"github.com/gasandov/academy-go-q32021/usecases"

	"github.com/labstack/echo/v4"
)

func CreateEchoRoutes(e *echo.Echo) *echo.Echo {
	healthHandler := controllers.NewHealthController()

	csvRepo := repositories.NewCSVRepo()

	pokemonService := usecases.NewPokemonService(csvRepo)
	pokemonHandler := controllers.NewPokemonController(pokemonService)

	consumerService := usecases.NewConsumerService(csvRepo)
	consumerHandler := controllers.NewConsumerController(consumerService)

	e.GET("/health-check", healthHandler.GetHealthCheck)

	e.GET("/pokemons", pokemonHandler.GetPokemons)
	e.GET("/pokemons/:id", pokemonHandler.GetPokemonById)

	e.GET("/consume", consumerHandler.ConsumeAPI)

	return e
}
