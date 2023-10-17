package helper

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Rep struct {
	Code    int         `json:"status_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"date"`
}

func Success(c *gin.Context, res *Rep) {
	c.JSON(http.StatusOK, res)

}

func Error(c *gin.Context, res *Rep) {

	c.JSON(http.StatusInternalServerError, res)
}

func JsonFail(c *gin.Context, res *Rep) {

	c.JSON(http.StatusInternalServerError, res)
}

// Try 异常捕获函数
func Try(fn func(), catchFn func(err interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catchFn(err)
		}
	}()
	fn()
}

func IsString(obj interface{}) bool {
	kind := reflect.ValueOf(obj).Kind()
	return kind == reflect.String
}
func IsInt(obj interface{}) bool {
	kind := reflect.ValueOf(obj).Kind()
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64
}
func IsFloat(obj interface{}) bool {
	kind := reflect.ValueOf(obj).Kind()
	return kind == reflect.Float32 || kind == reflect.Float64
}
func IsNumber(obj interface{}) bool {
	kind := reflect.ValueOf(obj).Kind()
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 || kind == reflect.Float32 || kind == reflect.Float64
}
func IsMap(obj interface{}) bool {
	kind := reflect.ValueOf(obj).Kind()
	return kind == reflect.Map
}
func IsSlice(obj interface{}) bool {
	kind := reflect.ValueOf(obj).Kind()
	return kind == reflect.Slice
}
func IsStruct(obj interface{}) bool {
	kind := reflect.ValueOf(obj).Kind()
	return kind == reflect.Struct
}
