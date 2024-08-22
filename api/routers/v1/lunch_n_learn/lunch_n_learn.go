package lunchnlearnroute

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	lnlfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/lunch-n-learn"
	dtos "github.com/ioet/ioet-lunch-n-learn-backend/api/dtos"
	lunchnlearn "github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/lunch_n_learn"
	lnlcreationusecase "github.com/ioet/ioet-lunch-n-learn-backend/core/src/use_cases/lunch_n_learn/create"
	lnllistingusecase "github.com/ioet/ioet-lunch-n-learn-backend/core/src/use_cases/lunch_n_learn/list/all"
	lnllistingbyidusecase "github.com/ioet/ioet-lunch-n-learn-backend/core/src/use_cases/lunch_n_learn/list/id"
	lnlmodificationusecase "github.com/ioet/ioet-lunch-n-learn-backend/core/src/use_cases/lunch_n_learn/update"
)

func Route(rg *gin.RouterGroup) {
	repository, err := lnlfirebaserepository.NewLnLRepository(context.Background())
	if err != nil {
		panic("Error initializing the Lunch-n-learn repository: " + err.Error())
	}

	rg.GET("/", func(c *gin.Context) {
		useCase := lnllistingusecase.NewLnLListingUseCase(*repository)
		lnls, err := useCase.Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error getting the Lunch-n-learns": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"lnls":    lnls,
			"count":   len(lnls),
			"message": "Lunch-n-learns retrieved successfully",
		})
	})

	rg.GET("/:id", func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid id format": err.Error()})
			return
		}

		useCase := lnllistingbyidusecase.NewLnLListingByIdUseCase(*repository)
		lnl, err := useCase.Execute(id.String())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error getting the Lunch-n-learn": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"lnl":     lnl,
			"message": "Lunch-n-learn retrieved successfully",
		})
	})

	rg.POST("/", func(c *gin.Context) {
		var newLnL dtos.LnLCreateIn
		if err := c.ShouldBindJSON(&newLnL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error reading Lunch-n-learn from request params": err.Error()})
			return
		}

		useCase := lnlcreationusecase.NewLnLCreationUseCase(*repository)
		_newlnl := lunchnlearn.New(newLnL.Name, newLnL.PresentationDate)
		createdLnL, err := useCase.Execute(_newlnl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error creating the Lunch-n-learn": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Lunch-n-learn created successfully",
			"lnl":     createdLnL,
		})
	})

	rg.PUT("/", func(c *gin.Context) {
		var requestLnL dtos.LnLUpdateIn
		if err := c.ShouldBindJSON(&requestLnL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error reading Lunch-n-learn from request params": err.Error()})
			return
		}

		useCase := lnlmodificationusecase.NewLnLModificationUseCase(*repository)
		updatedLnL, err := useCase.Execute(requestLnL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error updating the Lunch-n-learn": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Lunch-n-learn updated successfully",
			"lnl":     updatedLnL,
		})
	})

	rg.PUT("/assistant", func(c *gin.Context) {
		var requestLnL dtos.LnLAddAssistantIn
		if err := c.ShouldBindJSON(&requestLnL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error reading Lunch-n-learn from request params": err.Error()})
			return
		}

		useCase := lnlmodificationusecase.NewLnLAddAssistantUseCase(*repository)
		updatedLnL, err := useCase.Execute(requestLnL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error updating the Lunch-n-learn": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Lunch-n-learn updated successfully",
			"lnl":     updatedLnL,
		})
	})

	rg.PUT("/presenter", func(c *gin.Context) {
		var requestLnL dtos.LnLAddPresenterIn
		if err := c.ShouldBindJSON(&requestLnL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error reading Lunch-n-learn from request params": err.Error()})
			return
		}

		useCase := lnlmodificationusecase.NewLnLAddPresenterUseCase(*repository)
		updatedLnL, err := useCase.Execute(requestLnL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error updating the Lunch-n-learn": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Lunch-n-learn updated successfully",
			"lnl":     updatedLnL,
		})
	})
}
