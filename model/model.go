package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nmakro/best-price-api/db"
)

type DbModel struct {
	DbCon *db.DB
	// Products *Product
	// Category *Category
}

func (m *DbModel) Init() {
	m.DbCon = db.InitDb()
	m.DbCon.DbScope = m.DbCon.Db.NewScope(m.DbCon.Db)
	m.DbCon.Db.AutoMigrate(&Product{}, &Category{})

	// for i := 0; i < 200; i++ {
	// 	product := new(Products)
	// 	product.Title = fmt.Sprintf("title %d", i)
	// 	product.Category = fmt.Sprint("category %d", i%10)
	// 	product.Description = fmt.Sprintf("description %d", i)
	// 	product.Price = i * 3
	// 	m.DbCon.Db.Create(&product)

	// }

}

// func (p *Product) BeforeCreate(scope *gorm.Scope) error {
// 	return scope.SetColumn("ID", uuid.NewV4().String())
// }

// func (c *Category) BeforeCreate(scope *gorm.Scope) error {
// 	return scope.SetColumn("ID", uuid.NewV4().String())
// }
