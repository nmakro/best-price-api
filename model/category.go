package model

import (
	"time"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/nmakro/best-price-api/utils"
)

type Category struct {
	gorm.Model
	Title     string     `json:"title" gorm:"unique"`
	Position  int        `json:"position"`
	ImageURL  string     `json:"image_url"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (m *DbModel) CreateCategory(ctx *fiber.Ctx) {
	db := m.DbCon.Db

	category := new(Category)
	if err := ctx.BodyParser(category); err != nil {
		ctx.Status(503).Send(err)
		return
	}
	if category.Title == "" {
		ctx.Status(402).Send("Mandatory field Title is missing.\n")
	}
	db.Create(&category)
	ctx.JSON(category)
}

func (m *DbModel) GetCategories(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	var categories []Category
	p := utils.SetupPager(ctx, db)

	paginator := pagination.Paging(&p, &categories)

	meta := utils.CreateResponse(paginator)
	response := Response{Meta: meta, Categories: categories}
	ctx.JSON(response)
}

func (m *DbModel) GetCategory(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	id := ctx.Params("id")
	var category []Category
	db.Find(&category, id)
	ctx.JSON(&category)
}

func (m *DbModel) DeleteCategory(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	id := ctx.Params("id")
	if id == "" {
		ctx.Status(400).Send("Category id is missing\n")
	}
	var category Category
	//db.First(&category, id)
	if category == (Category{}) {
		ctx.Status(404).Send("Not found!\n")
	}
	db.Delete(category)

}
