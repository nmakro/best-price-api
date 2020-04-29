package model

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	//ID        int        `gorm:"type:int;primary_key"`
	Title    string `json:"title" gorm:"unique"`
	Position string `json:"position"`
	ImageURL string `json:"image_url"`
	// CreatedAt time.Time  `json:"created_at"`
	// UpdatedAt time.Time  `json:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	//Products  []Product  `json:"products" gorm:"foreignkey:ID"`
}

type MetaData struct {
	TotalRecords int `json:"total_records"`
	TotalPages   int `json:"total_pages"`
	Offset       int `json:"offset"`
	Limit        int `json:"limit"`
	Page         int `json:"page"`
	PrevPage     int `json:"prev_page"`
	NextPage     int `json:"next_page"`
}

type Response struct {
	Meta     MetaData   `json:"_meta"`
	Products []Product  `json:"products,omitempty"`
	Category []Category `json:"categories,omitempty"`
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
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &categories)

	fmt.Println(&paginator)
	fmt.Println(reflect.TypeOf(paginator))

	response := createResponse(paginator, categories, nil)
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
	var category Category
	db.First(&category, id)
	db.Delete(category)
	if category == (Category{}) {
		ctx.Status(503).Send("Not found!\n")
	}
}

func createResponse(paginator *pagination.Paginator, categories []Category, products []Product) (r Response) {

	metadata := MetaData{TotalRecords: paginator.TotalRecord, TotalPages: paginator.TotalPage, Offset: paginator.Offset,
		Limit: paginator.Limit, Page: paginator.Page, PrevPage: paginator.PrevPage, NextPage: paginator.NextPage}

	return Response{Meta: metadata, Category: categories, Products: products}
}
