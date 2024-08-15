package house_route

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/house"
)

// @Summary Get mocked houses
// @Description get all houses
// @ID get-all-houses
// @Accept  json
// @Produce  json
// @Success 200 {string} string  "ok"
// @Router /house [get]
func Route(rg *gin.RouterGroup) {
	rg.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"houses": generateMockHouses(),
			"count":  5,
		})
	})
}

func generateMockHouses() []house.House {
	mockHouses := make([]house.House, 5)

	for i := range mockHouses {
		mockHouses[i] = house.New("Mocked house "+strconv.Itoa(i), nil)
	}

	return mockHouses
}
