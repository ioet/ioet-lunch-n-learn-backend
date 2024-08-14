package events

import (
	"github.com/gin-gonic/gin"
	healthcheck_route "github.com/ioet/ioet-lunch-n-learn-backend/api/routers/v1/health_check"
	house_route "github.com/ioet/ioet-lunch-n-learn-backend/api/routers/v1/house"
	user_route "github.com/ioet/ioet-lunch-n-learn-backend/api/routers/v1/user"
)

// TODO: Modify base logic to register all routes dynamically
func RegisterRoutes(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	healthcheck_route.Route(v1.Group("/health_check"))
	user_route.Route(v1.Group("/user"))
	house_route.Route(v1.Group("/house"))

	// add another route
	// Example:
	// user.Route(v1.Group("/user"))
}
