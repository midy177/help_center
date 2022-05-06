package service

import (
	"github.com/labstack/echo/v4"
)

func UserRouter(user *echo.Group) {
	user.GET("/get_article", useGetArticle)
	user.GET("/match_article", useMatchArticle)
	user.GET("/get_category", useGetCategory)
}

func AdmRouter(adm *echo.Group) {
	adm.POST("/add_article", admAddArticle)
	adm.POST("/del_article", admDelArticle)
	adm.POST("/mod_article", admModArticle)
	adm.GET("/get_article", admGetArticle)
	adm.POST("/add_category", admAddCategory)
	adm.POST("/del_category", admDelCategory)
	adm.POST("/mod_category", admModCategory)
	adm.GET("/get_category", useGetCategory)
	adm.POST("/upload_img", admUploadImg)
}
func AuthRouter(user *echo.Group) {
	user.POST("/login", admLogin)
}
