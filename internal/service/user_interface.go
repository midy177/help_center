package service

import (
	"github.com/labstack/echo/v4"
	"help_center/internal/biz"
	"help_center/internal/data"
	"log"
	"net/http"
)

// useGetArticle doc
// @Tags UseApi
// @Summary 查询文章
// @Param body query data.ArticleQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=[]data.Article} "返回数据"
// @Router /api/get_article [GET]
func useGetArticle(c echo.Context) error {
	d := new(data.ArticleQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	d.Status = 2
	msg := biz.GetArticle(d, false)
	go StoreCache(c.Request().URL.Path, &msg)
	return c.JSON(http.StatusOK, &msg)
}

// useGetCategory doc
// @Tags UseApi
// @Summary 查询分类
// @Param body query data.CategoryQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=[]data.Category} "返回数据"
// @Router /api/get_category [GET]
func useGetCategory(c echo.Context) error {
	d := new(data.CategoryQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetCategory(d, false)
	go StoreCache(c.Request().URL.Path, &msg)
	return c.JSON(http.StatusOK, &msg)
}

// useMatchArticle
// @Tags UseApi
// @Summary title关键字查询文章(至多返回30条数据)
// @Param sub_str query string true "匹配数据"
// @Param lang query string true "语言"
// @Success 200 {object} biz.BaseJson{data=[]data.Article} "返回数据"
// @Router /api/match_article [GET]
func useMatchArticle(c echo.Context) error {
	lang := c.QueryParam("lang")
	subStr := c.QueryParam("sub_str")
	msg := biz.MatchArticle(subStr, lang)
	go StoreCache(c.Request().URL.Path, &msg)
	return c.JSON(http.StatusOK, &msg)
}
