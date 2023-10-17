package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func UserShow(c *gin.Context) {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	fmt.Println("user show")
	c.JSON(http.StatusOK, gin.H{
		"message": "usercontroller",
	})
}
