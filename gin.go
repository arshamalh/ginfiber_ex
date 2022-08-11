package main

import "github.com/gin-gonic/gin"

type ginApp struct {
	app *gin.Engine
}

// Make new instance of gin app
func NewGinApp() *ginApp {
	return &ginApp{
		app: gin.Default(),
	}
}

// Add routes to gin app
func (g *ginApp) RegisterRoutes() {
	g.app.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
}

// Start gin server
func (g *ginApp) Start() {
	g.app.Run(":80")
}
