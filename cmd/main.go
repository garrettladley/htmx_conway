package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/garrettladley/htmx_conway/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	app := fiber.New()

	app.Get("/:name?", func(c *fiber.Ctx) error {
		handler := adaptor.HTTPHandler(templ.Handler(views.Home(c.Params("name"))))
		return handler(c)
	})

	app.Use(NotFoundMiddleware)

	log.Fatal(app.Listen(":42069"))
}

func NotFoundMiddleware(c *fiber.Ctx) error {
	return Render(c, views.NotFound(), templ.WithStatus(http.StatusNotFound))
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
