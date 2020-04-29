package model

import (
	"fmt"
	"strconv"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	//ID          int        `gorm:"type:int;primary_key"`
	CategoryID  string `json:"category_id" gorm:"foreignkey:ID"`
	Title       string `json:"title"`
	ImageURL    string `json:"image_url"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	// CreatedAt   time.Time  `json:"created_at"`
	// UpdatedAt   time.Time  `json:"updated_at"`
	// DeletedAt   *time.Time `json:"deleted_at" sql:"index"`
}

func (m *DbModel) CreateProduct(ctx *fiber.Ctx) {
	db := m.DbCon.Db

	product := new(Product)
	if err := ctx.BodyParser(product); err != nil {
		ctx.Status(503).Send(err)
		return
	}
	db.Create(&product)
	ctx.JSON(product)
}

func (m *DbModel) GetProduct(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	id := ctx.Params("id")
	fmt.Println(id)
	var products []Product
	db.Find(&products, id)
	ctx.JSON(&products)
}

func (m *DbModel) GetProducts(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	var products []Product
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &products)

	//db.Find(&products)
	// if err := ctx.JSON(&products); err != nil {
	// 	fmt.Println(err)
	// }
	ctx.JSON(paginator)
}

func (m *DbModel) DeleteProduct(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	id := ctx.Params("id")
	var product Product
	db.First(&product, id)
	db.Delete(product)
	if product == (Product{}) {
		ctx.Status(503).Send("Not found!\n")
	}
}
