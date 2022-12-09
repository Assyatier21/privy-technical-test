package routes

import (
	"net/http"
	"privy/internal/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRoutes(handler api.Handler) *echo.Echo {
	e := echo.New()
	useMiddlewares(e)

	// CRUD User
	e.GET("/cakes", handler.GetListOfCakes)
	e.GET("/cakes/:id", handler.GetDetailsOfCake)
	e.POST("/cakes", handler.InsertCake)
	e.PATCH("/cakes/:id", handler.UpdateCake)
	e.DELETE("/cakes/:id", handler.DeleteCake)
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
