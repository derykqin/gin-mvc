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

	r.GET("/user/1", controller.UserShow)
	return r
}
