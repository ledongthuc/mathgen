package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/ledongthuc/mathgen/web/messages"
)

func setupWebRoutes(g *echo.Group) {
	g.Static("/", getAssetRootPath())
	g.GET("/", rootHandler)
	g.Static("/styles", getAssetStylePath())
	g.Static("/scripts", getAssetScriptPath())
	g.Static("/images", getAssetImagesPath())
}

func getAssetHTMLPath() string {
	return fmt.Sprintf("%s/html/*.html", getAssetPath())
}

func getAssetRootPath() string {
	return fmt.Sprintf("%s/root", getAssetPath())
}

func getAssetStylePath() string {
	return fmt.Sprintf("%s/styles", getAssetPath())
}

func getAssetScriptPath() string {
	return fmt.Sprintf("%s/scripts", getAssetPath())
}

func getAssetImagesPath() string {
	return fmt.Sprintf("%s/images", getAssetPath())
}

func getAssetPath() string {
	asset, exist := os.LookupEnv("ASSET_PATH")
	if !exist {
		asset = "assets"
	}
	return asset
}

func rootHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home", messages.GenerateWebMessage())
}
