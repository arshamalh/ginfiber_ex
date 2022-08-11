package main

import (
	"github.com/gofiber/fiber/v2"
)

type fiberApp struct {
	app *fiber.App
}

func NewFiberApp() *fiberApp {
	return &fiberApp{
		app: fiber.New(),
	}
}

func (g *fiberApp) RegisterRoutes() {
	g.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

func (g *fiberApp) Start() {
	g.app.Listen(":81")
}

func (g *fiberApp) Stop() {
	g.app.Shutdown()
}
