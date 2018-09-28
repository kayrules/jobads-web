package main

import (
	"html/template"
	"io"

	"./config"
	controller "./controller"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	e.Renderer = renderer

	e.Static("/assets", "assets")

	// Named routes
	e.GET("/", controller.Home).Name = "home"
	e.GET("/login", controller.Login).Name = "login"
	e.POST("/login", controller.LoginPost).Name = "login-post"
	e.GET("/cart", controller.Cart).Name = "cart"
	e.POST("/logout", controller.Logout).Name = "logout"

	e.Server.Addr = ":" + config.Port
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
