package main

import (
	"github.com/nmakro/best-price-api/server"
)

func main() {
	app := server.NewApp()
	app.Serve()

}
