package user_route

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/user"
)

// @Summary Get mocked users
// @Description get all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {string} string  "ok"
// @Router /user [get]
func Route(rg *gin.RouterGroup) {
	rg.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": generateMockUsers(),
			"count": 10,
		})
	})
}

func generateMockUsers() []user.User {
	mockUsers := make([]user.User, 10)

	for i := range mockUsers {
		mockUsers[i] = user.New("Mocked user "+strconv.Itoa(i), "Mocked email", nil, nil)
	}

	return mockUsers
}
