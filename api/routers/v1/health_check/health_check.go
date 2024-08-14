package healthcheck_route

import (
	"github.com/gin-gonic/gin"
)

func Route(rg *gin.RouterGroup) {
	rg.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ioet-lunch-n-learn-app": "ALIVE",
		})
	})
}
