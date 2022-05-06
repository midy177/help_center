package biz

import (
	"help_center/internal/data"
	"log"
	"strings"
)

// GetArticle 获取文章列表
func GetArticle(d *data.ArticleQuery, isAdm bool) *JsonFormat {
	if d.Page == 0 {
		d.Page = 1
	}
	if d.PageSize == 0 {
		d.PageSize = 10
	}
	num := d.ArticleCount()
	if num > 0 {
		list := d.ArticleSearch(isAdm)
		return &JsonFormat{Code: 1, Page: d.Page, PageSize: d.PageSize, PageNum: num/d.PageSize + 1, ArticleNum: num, Data: list}
	}
	return &JsonFormat{Code: 0, Page: d.Page, PageSize: d.PageSize, PageNum: 0, ArticleNum: num, Data: "当前参数获取到的文章数量为0"}
}

// MatchArticle 获取文章列表
func MatchArticle(subStr, lang string) *JsonFormat {
	if subStr != "" || !strings.Contains("cn,en", lang) {
		list, lenList := data.ArticleMatch(subStr, lang)
		return &JsonFormat{Code: 1, Page: 1, PageSize: lenList, PageNum: 1, ArticleNum: lenList, Data: list}
	}
	return &JsonFormat{Code: 0, Page: 0, PageSize: 0, PageNum: 0, ArticleNum: 0, Data: "当前参数获取到的文章数量为0"}
}

// AddArticle 添加文章
func AddArticle(d *data.Article) *BaseJson {
	err := data.ArticleAdd(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功添加文章"}
	}
}

// DelArticle 删除文章
func DelArticle(d *data.DelQuery) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	if !strings.Contains("cn,en", d.Lang) {
		return &BaseJson{Code: 0, Data: "参数 lang 值应为cn或en"}
	}
	err := data.ArticleDelete(d.Id, d.Lang)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功成功删除文章"}
	}
}

// ModArticle 修改文章
func ModArticle(d *data.Article) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	if !strings.Contains("cn,en", d.Lang) {
		return &BaseJson{Code: 0, Data: "参数 lang 值应为cn或en"}
	}
	err := data.ArticleUpdate(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功修改文章"}
	}
}
