package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nmakro/best-price-api/config"
)

type DB struct {
	Db      *gorm.DB
	DbScope *gorm.Scope
}

func InitDb() (con *DB) {
	var dbCon *gorm.DB
	var err error

	dbConfig := config.DBConfig{DbDriver: "mysql", DbUser: "best-price", Password: "123",
		URI: "@tcp(127.0.0.1:3306)/", DbName: "best_price", Options: "?charset=utf8&parseTime=True"}

	if dbCon, err = gorm.Open(dbConfig.DbDriver, dbConfig.DbUser+":"+dbConfig.Password+dbConfig.URI+dbConfig.DbName+dbConfig.Options); err != nil {
		fmt.Println(err)
		panic("Connection to MySql failed.")
	}

	fmt.Println("Succesfull connection to DB")
	return &DB{Db: dbCon}
}
