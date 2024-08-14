package healthcheck

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Route(rg *gin.RouterGroup) {
	rg.GET("/", func(c *gin.Context) {
		log.Println("Health check endpoint hit")
		c.JSON(200, gin.H{
			"ioet-lunch-n-learn-app": "ALIVE",
		})
	})
}
