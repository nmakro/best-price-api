package server

import (
	"github.com/gofiber/fiber"
	"github.com/nmakro/best-price-api/model"
)

type App struct {
	server *fiber.App
	model  *model.DbModel
}

func NewApp() (app *App) {
	fb := fiber.New()
	m := model.DbModel{}
	m.Init()
	return &App{server: fb, model: &m}
}

func (app *App) RouteHandler() {
	app.server.Get("best-price-api/v1/products", app.model.GetProducts)
	app.server.Get("best-price-api/v1/products/:id", app.model.GetProduct)
	app.server.Post("best-price-api/v1/products", app.model.CreateProduct)
	app.server.Delete("best-price-api/v1/products/:id", app.model.DeleteProduct)

}

func (app *App) Serve() {
	app.RouteHandler()
	app.server.Listen(3000)
}
