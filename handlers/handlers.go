package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	aproperty := os.Getenv("MY_PROPERTY")
	c.JSON(200, gin.H{
		"message": "pong with :" + aproperty,
	})
}
