package routes

import (
	"StarWarsBackEnd/api"
	"StarWarsBackEnd/api/middlewares"

	"github.com/labstack/echo"
)

// New ...
func New() *echo.Echo {

	e := echo.New()
	authGroup := e.Group("/api/v1")

	middlewares.SetAuthMiddleWares(authGroup)

	api.MainGroup(e)
	api.AuthGroup(authGroup)

	return e
}
