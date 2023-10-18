package main

import (
	"derykqin/gin-mvc/logs"
	"derykqin/gin-mvc/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	//初始化日志框架logrus
	logs.InitLogrus()
}

func main() {
	r := gin.Default()
	// r = LoadHtml(r)                   //设置html模板
	r = routes.SetGlobalMiddleware(r) //设置全局中间件
	r = routes.LoadRoute(r)           //设置路由
	r.Run()                           // 监听并在 0.0.0.0:8080 上启动服务
}

func LoadHtml(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("views/**/*")
	return r
}
