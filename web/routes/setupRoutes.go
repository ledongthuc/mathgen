package routes

import "github.com/labstack/echo"

// SetupRoutes will setup all routes in systems. Sub routs will be manage details by setupXXX function
func SetupRoutes(e *echo.Echo) {
	setupWebRoutes(e.Group(""))
	g := e.Group("/api")
	setupAdditionRoutes(g.Group("/addition"))
	setupSubtractionRoutes(g.Group("/subtraction"))
}
