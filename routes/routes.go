package routes

import (
	"github.com/gasandov/academy-go-q32021/controllers"
	"github.com/gasandov/academy-go-q32021/repositories"
	"github.com/gasandov/academy-go-q32021/usecases"

	"github.com/labstack/echo/v4"
)

func CreateEchoRoutes(e *echo.Echo) *echo.Echo {
	healthController := controllers.NewHealthController()

	fileManagerRepo := repositories.NewFileManagerRepo()
	pokemonService := usecases.NewPokemonService(fileManagerRepo)
	pokemonController := controllers.NewPokemonController(pokemonService)

	e.GET("/health-check", healthController.GetHealthCheck)

	e.GET("/api/pokemons/concurrent", pokemonController.GetPokemonsConcurrently)
	e.GET("/api/pokemons/source", pokemonController.GetPokemonsFromAPI)
	e.GET("/api/pokemons/:id", pokemonController.GetPokemonById)
	e.GET("/api/pokemons", pokemonController.GetPokemons)

	return e
}
