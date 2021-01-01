package routes

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
)

func setupWebRoutes(g *echo.Group) {
	g.Static("/", getAssetRootPath())
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
