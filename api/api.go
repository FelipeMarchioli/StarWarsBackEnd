package api

import (
	"StarWarsBackEnd/handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.GET("/health-check", handlers.HealthCheck)
}

func AuthGroup(a *echo.Group) {
	// AuthRoute
	a.GET("/planetas", handlers.GetPlanets)
	// a.GET("/planeta", handlers.GetPlanetById)
	// a.GET("/planeta", handlers.GetPlanetByName)
	// a.POST("/planeta", handlers.AddPlanet)
	// a.DELETE("/planeta", handlers.DeletePlanetById)
}
