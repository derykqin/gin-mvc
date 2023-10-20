package controller

import (
	"context"
	"derykqin/gin-mvc/core/cache"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type UserController struct{}

var UserCont UserController

func init() {
	UserCont = UserController{}
}

// Display a listing of the resource.
func (u UserController) Index(c *gin.Context) {

}

// Show the form for creating a new resource.
func (u UserController) Create(c *gin.Context) {

}

// Store a newly created resource in storage.
func (u UserController) Store(c *gin.Context) {

}

// Display the specified resource.
func (u UserController) Show(c *gin.Context) {
	ctx := context.Background()
	err := cache.RedisClient.Set(ctx, "u", "123456", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := cache.RedisClient.Get(ctx, "us").Result()
	if err == redis.Nil {
		fmt.Println("u does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("u", val)
	}
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	fmt.Println("user show")
	c.JSON(http.StatusOK, gin.H{
		"message": "usercontroller",
	})
}

// Show the form for editing the specified resource.
func (u UserController) Edit(c *gin.Context) {

}

// Update the specified resource in storage.
func (u UserController) Update(c *gin.Context) {

}

// Remove the specified resource from storage.
func (u UserController) Destroy(c *gin.Context) {

}
