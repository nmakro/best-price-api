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
	//app.server.Patch("best-price-api/v1/products/:id", app.model.UpdateProduct)
	app.server.Delete("best-price-api/v1/products/:id", app.model.DeleteProduct)

	app.server.Get("best-price-api/v1/categories", app.model.GetCategories)
	app.server.Get("best-price-api/v1/categories/:id", app.model.GetCategory)
	app.server.Post("best-price-api/v1/categories", app.model.CreateCategory)
	app.server.Delete("best-price-api/v1/categories", app.model.DeleteCategory)

}

func (app *App) Serve() {
	app.RouteHandler()
	app.server.Listen(3000)
}
