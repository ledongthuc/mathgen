package routes

import "github.com/labstack/echo"

func SetupRoutes(e *echo.Echo) {
	setupWebRoutes(e.Group(""))
	setupAdditionRoutes(e.Group("/addition"))
}
