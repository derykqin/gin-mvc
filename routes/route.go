package routes

import (
	"derykqin/gin-mvc/controller"
	"derykqin/gin-mvc/middlewares"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetGlobalMiddleware(r *gin.Engine) *gin.Engine {
	// 	r.Use(Mid1(), Mid2())

	return r
}
func LoadRoute(r *gin.Engine) *gin.Engine {
	r.GET("/ping", middlewares.Middleware1(), middlewares.Middleware2(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
		fmt.Println("it is controller")
	})

	v1 := r.Group("")
	v1.Use(middlewares.Logger()) //组路由中间件
	//资源路由
	v1.GET("/users", controller.UserCont.Index)
	v1.GET("/users/create", controller.UserCont.Create)
	v1.POST("users", controller.UserCont.Store)
	v1.GET("/users/:id", controller.UserCont.Show)
	v1.GET("/users/:id/edit", controller.UserCont.Edit)
	v1.PUT("/users/:id", controller.UserCont.Update)
	v1.DELETE("/users/:id", controller.UserCont.Destroy)
	return r
}
