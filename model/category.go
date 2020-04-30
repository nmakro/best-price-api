package model

import (
	"strconv"
	"time"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gofiber/fiber"
	"github.com/nmakro/best-price-api/utils"
)

type Category struct {
	//gorm.Model
	ID        uint       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
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
		ctx.Status(400).Send(err)
		return
	}
	if category.Title == "" {
		ctx.Status(400).Send("Mandatory field Title is missing.\n")
	}
	if err := db.Create(&category).Error; err != nil {
		ctx.Status(500).Send(err.Error())
	}
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
	var category Category
	db.Find(&category, id)
	if category == (Category{}) {
		ctx.Status(404).Send("Category not found!\n")
		return
	}
	ctx.JSON(&category)
}

func (m *DbModel) UpdateCategory(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	id := ctx.Params("id")

	category := new(Category)
	if err := ctx.BodyParser(category); err != nil {
		ctx.Status(422).Send(err)
		return
	}
	categoryInDb := new(Category)

	db.First(&categoryInDb, id)
	if categoryInDb.ID == 0 {
		ctx.Status(404).Send("Category not found!\n")
	}

	category.ID = categoryInDb.ID

	if err := db.Debug().Model(&categoryInDb).Updates(category).Error; err != nil {
		ctx.Status(500).Send("Cannot update category.\n")
	}
}

func (m *DbModel) DeleteCategory(ctx *fiber.Ctx) {
	db := m.DbCon.Db
	id := ctx.Params("id")
	var category Category
	a_id, err := strconv.Atoi(id)
	if err == nil {
		category.ID = uint(a_id)
	} else {
		ctx.Status(400).Send("Unable to parse category id!\n")
		return
	}
	if err := db.Delete(category).Error; err != nil {
		ctx.Status(500).Send("Delete failed!\n")
	}
	db.Model(category).Unscoped().Debug().Update("title", nil)
	ctx.Status(204)
}
