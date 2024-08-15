package events

import (
	"github.com/gin-gonic/gin"
	healthcheckroute "github.com/ioet/ioet-lunch-n-learn-backend/api/routers/v1/health_check"
	houseroute "github.com/ioet/ioet-lunch-n-learn-backend/api/routers/v1/house"
	lunchnlearnroute "github.com/ioet/ioet-lunch-n-learn-backend/api/routers/v1/lunch_n_learn"
	userroute "github.com/ioet/ioet-lunch-n-learn-backend/api/routers/v1/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// TODO: Modify base logic to register all routes dynamically
func RegisterRoutes(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	healthcheckroute.Route(v1.Group("/health_check"))
	userroute.Route(v1.Group("/user"))
	houseroute.Route(v1.Group("/house"))
	lunchnlearnroute.Route(v1.Group("/lunch_n_learn"))

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// add another route
	// Example:
	// user.Route(v1.Group("/user"))
}
