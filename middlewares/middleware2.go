package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before methon,it is the middleware2")
		c.Next()
		fmt.Println("after methon,it is the middleware2")
	}
}
