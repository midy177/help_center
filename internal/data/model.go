package data

import (
	"time"
)

type Article struct {
	Id       int64     `json:"id"`
	Lang     string    `json:"lang"`
	Status   int       `json:"status"`
	CateId   int64     `json:"cate_id"`
	Title    string    `json:"title"`
	Summary  string    `json:"summary"`
	Markdown string    `json:"markdown"`
	RichText string    `json:"rich_text"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Uri      string    `json:"uri"`
}

type Category struct {
	Id       int64  `json:"id"`
	Lang     string `json:"lang"`
	Name     string `json:"name"`
	Intro    string `json:"intro"`
	ParentId int64  `json:"parent_id"`
	//List     string `json:"list"`
}

type CategoryRelation struct {
	Id       int64 `json:"id"`
	ParentId int64 `json:"parent_id"`
}

type ArticleQuery struct {
	Id       int64  `json:"id" query:"id"`
	Status   int    `json:"status" query:"status"`
	CateId   int64  `json:"cate_id" query:"cate_id"`
	Lang     string `json:"lang" query:"lang"`
	Page     int    `json:"page" query:"page"`
	PageSize int    `json:"page_size" query:"page_size"`
}
type CategoryQuery struct {
	Id       int64  `json:"id" query:"id"`
	Lang     string `json:"lang" query:"lang"`
	ParentId int64  `json:"parent_id" query:"parent_id"`
	//Page     int    `json:"page" query:"page"`
	//PageSize int    `json:"page_size" query:"page_size"`
}
type DelQuery struct {
	Id   int64  `json:"id" query:"id"`
	Lang string `json:"lang" query:"lang"`
}
