package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ledongthuc/mathgen"
	"github.com/ledongthuc/mathgen/web/messages"
)

func setupAdditionRoutes(g *echo.Group) {
	g.POST("/generate", func(c echo.Context) error {
		var request messages.AdditionRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if valid, err := request.Valid(); !valid {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		result := mathgen.AddIntegerN(request.NumberOfAddend, request.MaxSum)
		return c.JSON(http.StatusOK, messages.AdditionResultFromModel(result))
	})
}
