package main

import (
	"derykqin/gin-mvc/routes"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	//初始化日志框架logrus
	log.SetFormatter(&log.JSONFormatter{}) //text formatter or json formatter
	// log.SetReportCaller(true)//添加记录日志方法
	log.SetLevel(log.InfoLevel) //info level
	writer, err := os.OpenFile("./logs/log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, writer)) //输出多个通道

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