package utils

import (
	"strconv"
	"strings"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type MetaData struct {
	TotalRecords int `json:"total_records"`
	TotalPages   int `json:"total_pages"`
	PerPage      int `json:"per_page"`
	Page         int `json:"page"`
	PrevPage     int `json:"prev_page"`
	NextPage     int `json:"next_page"`
}

func CreateResponse(paginator *pagination.Paginator) (m MetaData) {

	metadata := MetaData{TotalRecords: paginator.TotalRecord, TotalPages: paginator.TotalPage,
		PerPage: paginator.Limit, Page: paginator.Page, PrevPage: paginator.PrevPage, NextPage: paginator.NextPage}

	return metadata
}

func SetupPager(ctx *fiber.Ctx, db *gorm.DB) (par pagination.Param) {
	var orderBy []string
	page, _ := strconv.Atoi(ctx.Query("page"))
	per_page, _ := strconv.Atoi(ctx.Query("per_page"))
	sort := strings.SplitN(ctx.Query("sort"), ":", 2)

	orderBy = []string{"id desc"}

	if len(sort) == 2 {
		orderBy = []string{sort[0] + " " + sort[1]}
	}
	var p pagination.Param
	p.DB = db
	p.Limit = per_page
	p.Page = page
	p.OrderBy = orderBy
	p.ShowSQL = true
	return p
}
