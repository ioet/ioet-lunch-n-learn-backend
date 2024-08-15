package healthcheck_route

import (
	"github.com/gin-gonic/gin"
)

// @Summary Health Check
// @Description health check
// @ID health-check
// @Produce  json
// @Success 200 {string} string  "ok"
// @Router /health_check [get]
func Route(rg *gin.RouterGroup) {
	rg.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ioet-lunch-n-learn-app": "ALIVE",
		})
	})
}
