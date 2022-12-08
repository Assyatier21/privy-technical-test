package routes

import (
	"net/http"
	"privy/internal/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRoutes() *echo.Echo {
	e := echo.New()
	useMiddlewares(e)

	// CRUD User
	e.GET("/cakes", api.GetListOfCakes)
	e.GET("/cakes/:id", api.GetDetailsOfCake)
	e.POST("/cakes", api.InsertCake)
	e.PATCH("/cakes/:id", api.UpdateCake)
	e.DELETE("/cakes/:id", api.DeleteCake)
	return e
}

func useMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))
}
