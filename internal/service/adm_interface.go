package service

import (
	"github.com/labstack/echo/v4"
	"help_center/internal/biz"
	"help_center/internal/data"
	"log"
	"net/http"
	"strings"
)

// admAddArticle doc
// @Tags Article-文章
// @Summary 添加文章
// @Param token header string true "token"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/upload_img [POST]
func admUploadImg(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusOK, biz.BaseJson{Code: 0, Data: err.Error()})
	}
	contentType := file.Header.Get("Content-Type")
	if !strings.Contains(contentType, "image") {
		return c.JSON(http.StatusOK, biz.BaseJson{Code: 0, Data: "请选择图片文件"})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusOK, biz.BaseJson{Code: 0, Data: "请选择图片文件"})
	}
	msg := biz.UploadImage(src, file.Filename)
	return c.JSON(http.StatusOK, &msg)
}

// admGetArticle doc
// @Tags Article-文章
// @Summary 查询文章
// @Param token header string true "token"
// @Param body query data.ArticleQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=[]data.Article} "返回数据"
// @Router /adm/get_article [GET]
func admGetArticle(c echo.Context) error {
	d := new(data.ArticleQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	d.Status = 0
	msg := biz.GetArticle(d, true)
	return c.JSON(http.StatusOK, &msg)
}

// admAddArticle doc
// @Tags Article-文章
// @Summary 添加文章
// @Param token header string true "token"
// @Param body body data.Article true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/add_article [POST]
func admAddArticle(c echo.Context) error {
	d := new(data.Article)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	//d.Status = 1
	msg := biz.AddArticle(d)
	return c.JSON(http.StatusOK, &msg)
}

// admDelArticle doc
// @Tags Article-文章
// @Summary 删除文章
// @Param token header string true "token"
// @Param body body data.DelQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/del_article [POST]
func admDelArticle(c echo.Context) error {
	d := new(data.DelQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.DelArticle(d)
	return c.JSON(http.StatusOK, &msg)
}

// admModArticle doc
// @Tags Article-文章
// @Summary 修改文章
// @Param token header string true "token"
// @Param body body data.Category true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/mod_article [POST]
func admModArticle(c echo.Context) error {
	d := new(data.Article)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.ModArticle(d)
	return c.JSON(http.StatusOK, &msg)
}

// admAddCategory doc
// @Tags Category-分类
// @Summary 增加分类
// @Param token header string true "token"
// @Param body body data.Category true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/add_category [POST]
func admAddCategory(c echo.Context) error {
	d := new(data.Category)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.AddCategory(d)
	return c.JSON(http.StatusOK, &msg)
}

// admDelCategory doc
// @Tags Category-分类
// @Summary 删除分类
// @Param token header string true "token"
// @Param body body data.DelQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/del_category [POST]
func admDelCategory(c echo.Context) error {
	d := new(data.DelQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.DelCategory(d)
	return c.JSON(http.StatusOK, &msg)
}

// admModCategory doc
// @Tags Category-分类
// @Summary 修改分类
// @Param token header string true "token"
// @Param body body data.Category true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/mod_category [post]
func admModCategory(c echo.Context) error {
	d := new(data.Category)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.ModCategory(d)
	return c.JSON(http.StatusOK, &msg)
}

// admLogin doc
// @Tags auth-登陆认证
// @Summary 登陆
// @Param body body biz.LoginData true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /auth/login [post]
func admLogin(c echo.Context) error {
	loginData := new(biz.LoginData)
	err := c.Bind(loginData)
	if err != nil {
		log.Println(err.Error())
	}
	code, msg := biz.AdminLogin(loginData)
	return c.JSON(code, &msg)
}
