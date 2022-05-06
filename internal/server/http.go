package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "help_center/docs"
	"help_center/internal/conf"
	"help_center/internal/service"
	"log"
)

// RunApp 入口
// @Title support's 接口文档
// @Version 1.0
// @Description token传递方式包括[get/post]token、[header]Authorization
// @Description /api/* 公共访问
// @Description /adm/* 必须传入 token
// @Host 127.0.0.1:8080
// @BasePath /
func RunApp() {
	engine := echo.New()
	engine.Pre(service.CacheMidPre) // 实例化echo
	engine.Use(service.MidRecover)  // 恢复 日志记录
	//engine.HTTPErrorHandler = HTTPErrorHandler         // 自定义错误处理
	//engine.Pre(service.Swagger)
	engine.Debug = false // 运行模式 - echo框架好像没怎么使用这个
	//engine.GET("/swagger/index.html", echoSwagger.WrapHandler)
	engine.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Index:  "index.html",
		HTML5:  true,
		Browse: false,
	}))
	go func() {
		swag := echo.New()
		swag.GET("/swagger/*", echoSwagger.WrapHandler)
		err := swag.Start("0.0.0.0:8099")
		if err != nil {
			log.Fatalf("run error :%s", err.Error())
		}
	}()
	user := engine.Group("/api")
	service.UserRouter(user)
	auth := engine.Group("/auth")
	service.AuthRouter(auth)
	adm := engine.Group("/adm")
	secret := conf.GetCfg().Jwt.Secret
	adm.Use(middleware.JWT([]byte(secret)))
	service.AdmRouter(adm)
	err := engine.Start(conf.GetCfg().Echo.Listen)
	if err != nil {
		log.Fatalf("run error :%s", err.Error())
	}
}
