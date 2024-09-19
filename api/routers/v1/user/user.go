package userroute

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	dtos "github.com/ioet/ioet-lunch-n-learn-backend/api/dtos"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/user"
	usercreationusecase "github.com/ioet/ioet-lunch-n-learn-backend/core/src/useCases/user/create"
	userlistingusecase "github.com/ioet/ioet-lunch-n-learn-backend/core/src/useCases/user/list/all"
	userlistingbyidusecase "github.com/ioet/ioet-lunch-n-learn-backend/core/src/useCases/user/list/id"
	usermodificationusecase "github.com/ioet/ioet-lunch-n-learn-backend/core/src/useCases/user/update"
	repositoryfactories "github.com/ioet/ioet-lunch-n-learn-backend/factories/repositories"
)

func Route(rg *gin.RouterGroup) {
	repository := repositoryfactories.UserFirebaseRepository()

	rg.GET("/", func(c *gin.Context) {
		useCase := userlistingusecase.NewUserListingUseCase(*repository)
		users, err := useCase.Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error getting the Users": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"users":   users,
			"count":   len(users),
			"message": "Users retrieved successfully",
		})
	})

	rg.GET("/:id", func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid id format": err.Error()})
			return
		}

		useCase := userlistingbyidusecase.NewUserListingByIdUseCase(*repository)
		user, err := useCase.Execute(id.String())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error getting the User": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user":    user,
			"message": "User retrieved successfully",
		})
	})

	rg.POST("/", func(c *gin.Context) {
		var newUser dtos.UserCreateIn
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error reading User from request params": err.Error()})
			return
		}

		useCase := usercreationusecase.NewUserCreationUseCase(*repository)
		_newuser := user.New(newUser.Name, newUser.Email, newUser.PhotoURL, newUser.HouseID)
		createdUser, err := useCase.Execute(_newuser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error creating the User": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "House created successfully",
			"user":    createdUser,
		})
	})

	rg.PUT("/house", func(c *gin.Context) {
		var newUser dtos.UserChangeHouseIn
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error reading User from request params": err.Error()})
			return
		}

		useCase := usermodificationusecase.NewUserHouseChangeUseCase(*repository)
		updatedUser, err := useCase.Execute(newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error updating the User": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "House updated successfully",
			"user":    updatedUser,
		})
	})

	rg.PUT("/points", func(c *gin.Context) {
		var newUser dtos.UserAddPointsIn
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error reading User from request params": err.Error()})
			return
		}

		useCase := usermodificationusecase.NewUserAddPointsUseCase(*repository)
		updatedUser, err := useCase.Execute(newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error updating the User": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Points updated successfully",
			"user":    updatedUser,
		})
	})
}
