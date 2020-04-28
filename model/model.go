package model

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nmakro/best-price-api/db"
)

type Products struct {
	gorm.Model
	//ID       uint   `json:"id,omitempty" gorm:"column:id"`
	Category    string `json:"category,omitempty" gorm:"column:category"`
	Title       string `json:"title,omitempty  gorm:"column:title"`
	ImgURL      string `json:"imgUrl,omitempty  gorm:"column:imgURL"`
	Price       int    `json:"price,omitempty  gorm:"column:price"`
	Description string `json:"description,omitempty  gorm:"column:description"`
	// CreatedAt   time.Time  `json:"created_at,omitempty  gorm:"column:created_at"`
	// UpdatedAt   time.Time  `json:"updated_at,omitempty  gorm:"column:updated_at"`
	// DeletedAt   *time.Time `json:"updated_at,omitempty"  gorm:"column:updated_at sql:"index"`
}

type DbModel struct {
	//Products Products
	DbCon *db.DB
}

func (m *DbModel) Init() {
	m.DbCon = db.InitDb()
	m.DbCon.Db.AutoMigrate(&Products{})

	// for i := 0; i < 200; i++ {
	// 	product := new(Products)
	// 	product.Title = fmt.Sprintf("title %d", i)
	// 	product.Category = fmt.Sprint("category %d", i%10)
	// 	product.Description = fmt.Sprintf("description %d", i)
	// 	product.Price = i * 3
	// 	m.DbCon.Db.Create(&product)

	// }

}

func (m *DbModel) GetProduct(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	id := ctx.Params("id")
	var products []Products
	db.Find(&products, id)
	ctx.JSON(&products)
}

func (m *DbModel) GetProducts(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	page := ctx.Query("page")
	offset := ctx.Query("offset")
	fmt.Println(page)
	fmt.Println(offset)
	var products []Products
	db.Find(&products)
	if err := ctx.JSON(&products); err != nil {
		fmt.Println(err)
	}
	ctx.JSON(products)
}

func (m *DbModel) DeleteProduct(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	id := ctx.Params("id")
	var product Products
	db.First(&product, id)
	db.Delete(product)
	if product.ID == 0 {
		ctx.Status(503).Send("Not found!")
	}
}

func (m *DbModel) CreateProduct(ctx *fiber.Ctx) {
	db := m.DbCon.Db

	product := new(Products)
	if err := ctx.BodyParser(product); err != nil {
		ctx.Status(503).Send(err)
		return
	}
	db.Create(&product)
	ctx.JSON(product)
}
