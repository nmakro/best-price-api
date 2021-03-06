package server

import (
	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
	"github.com/nmakro/best-price-api/config"
	"github.com/nmakro/best-price-api/model"
)

type App struct {
	server *fiber.App
	model  *model.DbModel
	auth   func(*fiber.Ctx)
}

func NewApp() (app *App) {
	fb := fiber.New()
	authConfig := config.AuthConfig{Username: "admin", Password: "simple_pswd"}
	auth := basicauth.New(basicauth.Config{
		Users: map[string]string{
			authConfig.Username: authConfig.Password,
		},
	})
	m := model.DbModel{}
	m.Init()
	return &App{server: fb, model: &m, auth: auth}
}

func (app *App) RouteHandler() {

	app.server.Get("best-price-api/v1/products", app.model.GetProducts)
	app.server.Get("best-price-api/v1/products/:id", app.model.GetProduct)
	app.server.Post("best-price-api/v1/products", app.auth, app.model.CreateProduct)
	app.server.Patch("best-price-api/v1/products/:id", app.auth, app.model.UpdateProduct)
	app.server.Delete("best-price-api/v1/products/:id", app.auth, app.model.DeleteProduct)

	app.server.Get("best-price-api/v1/categories", app.model.GetCategories)
	app.server.Get("best-price-api/v1/categories/:id", app.model.GetCategory)
	app.server.Post("best-price-api/v1/categories", app.auth, app.model.CreateCategory)
	app.server.Patch("best-price-api/v1/categories/:id", app.auth, app.model.UpdateCategory)
	app.server.Delete("best-price-api/v1/categories/:id", app.auth, app.model.DeleteCategory)

}

func (app *App) Serve() {
	app.RouteHandler()
	app.server.Listen(3000)
}
