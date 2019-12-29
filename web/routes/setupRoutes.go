package routes

import "github.com/labstack/echo"

func SetupRoutes(e *echo.Echo) {
	setupWebRoutes(e.Group(""))
	g := e.Group("/api")
	setupAdditionRoutes(g.Group("/addition"))
}
