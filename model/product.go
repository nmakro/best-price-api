package model

import (
	"strconv"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/nmakro/best-price-api/utils"
)

type Product struct {
	gorm.Model
	//ID          int        `gorm:"type:int;primary_key"`
	Category    Category `json:"-" gorm:"foreignkey:CategoryID; AssociationForeignKey:Refer"`
	CategoryID  uint     `json:"category_id"`
	Title       string   `json:"title"`
	ImageURL    string   `json:"image_url"`
	Price       float32  `json:"price"`
	Description string   `json:"description"`
	// DeletedAt   *time.Time `json:"deleted_at" sql:"index"`
	// CreatedAt   time.Time  `json:"created_at"`
	// UpdatedAt   time.Time  `json:"updated_at"`
}

func (m *DbModel) CreateProduct(ctx *fiber.Ctx) {
	db := m.DbCon.Db

	product := new(Product)
	if err := ctx.BodyParser(product); err != nil {
		ctx.Status(503).Send(err)
		return
	}
	if err := db.Create(&product).Error; err != nil {
		ctx.Status(500).Send("Product insertion failed\n.")
		return
	}
	ctx.JSON(product)
}

func (m *DbModel) GetProduct(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	id := ctx.Params("id")
	var product Product
	db.Find(&product, id)
	if product == (Product{}) {
		ctx.Status(404).Send("Product not found!\n")
		return
	}
	ctx.JSON(&product)
}

func (m *DbModel) GetProducts(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	var products []Product

	p := utils.SetupPager(ctx, db)

	paginator := pagination.Paging(&p, &products)
	meta := utils.CreateResponse(paginator)
	response := Response{Meta: meta, Products: products}

	ctx.JSON(response)
}

func (m *DbModel) UpdateProduct(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	id := ctx.Params("id")
	productInDb := new(Product)

	db.First(&productInDb, id)
	if productInDb == &(Product{}) {
		ctx.Status(404).Send("Product not found!\n")
	}
	product := new(Product)
	if err := ctx.BodyParser(product); err != nil {
		ctx.Status(422).Send(err)
		return
	}
	product.ID = productInDb.ID

	if err := db.Debug().Model(&productInDb).Updates(product).Error; err != nil {
		ctx.Status(500).Send("Cannot update product.\n")
	}
}

func (m *DbModel) DeleteProduct(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	id := ctx.Params("id")
	var product Product
	a_id, err := strconv.Atoi(id)
	if err == nil {
		product.ID = uint(a_id)
	} else {
		ctx.Status(400).Send("Unable to parse product id!\n")
		return
	}
	if err := db.Delete(product).Error; err != nil {
		ctx.Status(500).Send("Delete failed!\n")
	}
	ctx.Status(204)
}
