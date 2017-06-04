package ui

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	homeTemplate = "home"
)

// Template implements echo.Renderer
type Template struct {
	templates *template.Template
}

// Render implements echo.Renderer
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func healthz(c echo.Context) error {
	return nil
}

func home(c echo.Context) error {
	return c.Render(http.StatusOK, homeTemplate, nil)
}

// Start starts the UI server
func Start(port string) {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("assets/templates/*.html")),
	}
	e.Renderer = t
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./assets")
	e.GET("/", home)
	e.GET("/healthz", healthz)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
