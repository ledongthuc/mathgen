package routes

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
)

func setupWebRoutes(g *echo.Group) {
	g.Static("/", getAssetHtmlPath())
	g.Static("/styles", getAssetStylePath())
	g.Static("/scripts", getAssetScriptPath())
	g.Static("/images", getAssetImagesPath())
}

func getAssetHtmlPath() string {
	return fmt.Sprintf("%s/html", getAssetPath())
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
