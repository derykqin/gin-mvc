package middlewares

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Logger 记录每次请求的请求信息和响应信息
func Logger(c *gin.Context) {
	// 请求前
	t := time.Now()
	reqPath := c.Request.URL.Path
	// reqId := c.GetString(constant.RequestId)
	method := c.Request.Method
	ip := c.ClientIP()
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		requestBody = []byte{}
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

	log.WithFields(log.Fields{
		// "reqid":  reqId,
		"host":   ip,
		"path":   reqPath,
		"method": method,
		"body":   string(requestBody),
	}).Info(fmt.Sprintf("%s %s start", method, reqPath))
	c.Next()
	// 请求后
	latency := time.Since(t)
	log.WithFields(log.Fields{
		// "reqid": reqId,
		"host": ip,
		"path": reqPath,
		"cost": latency,
	}).Info(fmt.Sprintf("%s %s end", method, reqPath))
}
