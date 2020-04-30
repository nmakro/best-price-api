package model

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nmakro/best-price-api/db"
	"github.com/nmakro/best-price-api/utils"
)

type DbModel struct {
	DbCon *db.DB
}

type Response struct {
	Meta       utils.MetaData `json:"_meta"`
	Products   []Product      `json:"products,omitempty"`
	Categories []Category     `json:"categories,omitempty"`
}

func (m *DbModel) Init() {
	m.DbCon = db.InitDb()
	m.DbCon.DbScope = m.DbCon.Db.NewScope(m.DbCon.Db)
	m.DbCon.Db.AutoMigrate(&Product{}, &Category{})
	m.DbCon.Db.Model(&Product{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	fmt.Println("Db migrated!")

	// categoriesArray := []string{"books", "tech", "clothes", "music", "sports"}
	// for i := 0; i < 5; i++ {
	// 	category := new(Category)
	// 	category.Title = categoriesArray[i]
	// 	category.Position = i
	// 	category.ImageURL = fmt.Sprintf("http://image-url.com/%d", i)
	// 	m.DbCon.Db.Create(&category)

	// }
	// for i := 0; i < 200; i++ {
	// 	product := new(Product)
	// 	product.Title = fmt.Sprintf("title %d", i)
	// 	product.CategoryID = uint((i % 5) + 1)
	// 	product.Description = fmt.Sprintf("description %d", i)
	// 	product.Price = float32(i) * 3.0
	// 	m.DbCon.Db.Create(&product)

	// }

}

// func (p *Product) BeforeCreate(scope *gorm.Scope) error {
// 	return scope.SetColumn("ID", uuid.NewV4().String())
// }

// func (c *Category) BeforeCreate(scope *gorm.Scope) error {
// 	return scope.SetColumn("ID", uuid.NewV4().String())
// }
