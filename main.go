package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ioet/ioet-lunch-n-learn-backend/api/config"
	"github.com/ioet/ioet-lunch-n-learn-backend/api/events"
)

func main() {
	engine := gin.Default()
	engine.Use(gin.Logger())

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	events.RegisterRoutes(engine)

	engine.Run(":" + config.APP_PORT)
}
