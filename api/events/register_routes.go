package events

import (
	"github.com/gin-gonic/gin"
	healthcheck "github.com/ioet/ioet-lunch-n-learn-backend/api/routers/v1"
)

// TODO: Modify base logic to register all routes dynamically
func RegisterRoutes(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	healthcheck.Route(v1.Group("/health_check"))

	// add another route
	// Example:
	// user.Route(v1.Group("/user"))
}
