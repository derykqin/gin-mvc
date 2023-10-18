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

func UserShow(c *gin.Context) {
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
