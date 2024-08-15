package events

import (
	"github.com/gin-gonic/gin"
	healthcheck_route "github.com/ioet/ioet-lunch-n-learn-backend/api/routers/v1/health_check"
	house_route "github.com/ioet/ioet-lunch-n-learn-backend/api/routers/v1/house"
	lunch_n_learn_route "github.com/ioet/ioet-lunch-n-learn-backend/api/routers/v1/lunch_n_learn"
	user_route "github.com/ioet/ioet-lunch-n-learn-backend/api/routers/v1/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// TODO: Modify base logic to register all routes dynamically
func RegisterRoutes(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	healthcheck_route.Route(v1.Group("/health_check"))
	user_route.Route(v1.Group("/user"))
	house_route.Route(v1.Group("/house"))
	lunch_n_learn_route.Route(v1.Group("/lunch_n_learn"))

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// add another route
	// Example:
	// user.Route(v1.Group("/user"))
}
