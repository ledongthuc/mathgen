package routes

import "github.com/labstack/echo"

func SetupRoutes(e *echo.Echo) {
	setupAdditionRoutes(e.Group("/addition"))
}
