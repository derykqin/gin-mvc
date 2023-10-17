package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Middleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		// fmt.Println(config.Get("a1.b1.v1"))
		fmt.Println("before methon,it is the middleware1")
		c.Next()
		fmt.Println("after methon,it is the middleware1")
	}
}
