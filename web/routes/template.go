package routes

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func GetTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob(getAssetHTMLPath())),
	}
}
