package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ioet/ioet-lunch-n-learn-backend/api/config"
	"github.com/ioet/ioet-lunch-n-learn-backend/api/events"
	"github.com/ioet/ioet-lunch-n-learn-backend/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	docs.SwaggerInfo.Title = "Lunch & Learn App"
	docs.SwaggerInfo.Description = "IOET Lunch & Learn App Backend"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5050"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	engine := gin.Default()
	engine.Use(gin.Logger())

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	events.RegisterRoutes(engine)
	err := engine.Run("localhost:" + config.AppPort)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
