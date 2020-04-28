package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	Db *gorm.DB
}

func InitDb() (con *DB) {
	var dbCon *gorm.DB
	var err error
	if dbCon, err = gorm.Open("mysql", "best-price:123@tcp(127.0.0.1:3306)/best_price?charset=utf8&parseTime=True"); err != nil {
		fmt.Println(err)
		panic("Connection to MySql failed.")
	}

	fmt.Println("Succesfull connection to DB")

	//dbCon.AutoMigrate(&model.Products{})

	fmt.Println("Db migrated!")
	return &DB{Db: dbCon}
}
