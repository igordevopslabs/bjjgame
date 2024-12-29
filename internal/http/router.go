package httpcontroller

import (
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "bjjgame app is healthy.",
	})
}

func Ready(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "bjjgame app is ready to go.",
	})
}
