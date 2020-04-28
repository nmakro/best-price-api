package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nmakro/best-price-api/server"
)

func main() {
	app := server.NewApp()
	app.Serve()

}
