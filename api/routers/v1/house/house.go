package houseroute

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	dtos "github.com/ioet/ioet-lunch-n-learn-backend/api/dtos"
	house "github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/house"
	housecreationusecase "github.com/ioet/ioet-lunch-n-learn-backend/core/src/useCases/house/create"
	houselistingusecase "github.com/ioet/ioet-lunch-n-learn-backend/core/src/useCases/house/list/all"
	houselistingbyidusecase "github.com/ioet/ioet-lunch-n-learn-backend/core/src/useCases/house/list/id"
	orchestratorfactories "github.com/ioet/ioet-lunch-n-learn-backend/factories/orchestrators"
	repositoryfactories "github.com/ioet/ioet-lunch-n-learn-backend/factories/repositories"
)

func Route(rg *gin.RouterGroup) {
	repository := repositoryfactories.HouseFirebaseRepository()
	orchestrator := orchestratorfactories.HouseFirebaseOrchestrator()

	rg.GET("/", func(c *gin.Context) {
		useCase := houselistingusecase.NewHouseListingUseCase(*orchestrator)
		houses, err := useCase.Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error getting the Houses": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"houses":  houses,
			"count":   len(houses),
			"message": "Houses retrieved successfully",
		})
	})

	rg.GET("/:id", func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid id format": err.Error()})
			return
		}

		useCase := houselistingbyidusecase.NewHouseListingByIdUseCase(*orchestrator)
		house, err := useCase.Execute(id.String())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error getting the House": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"house":   house,
			"message": "House retrieved successfully",
		})
	})

	rg.POST("/", func(c *gin.Context) {
		var newHouse dtos.HouseCreateIn
		if err := c.ShouldBindJSON(&newHouse); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error reading House from request params": err.Error()})
			return
		}

		captainID, err := uuid.Parse(newHouse.CaptainID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid captainId format": err.Error()})
			return
		}

		useCase := housecreationusecase.NewHouseCreationUseCase(*repository)
		_newhouse := house.New(newHouse.Name, captainID.String())
		createdHouse, err := useCase.Execute(_newhouse)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error creating the House": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "House created successfully",
			"house":   createdHouse,
		})
	})
}
