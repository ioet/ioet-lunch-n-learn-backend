package lunchnlearnroute

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	lunchnlearn "github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/lunch_n_learn"
)

// @Summary Get mocked lunch n learns
// @Description get all lunch n learns
// @ID get-all-lunch-n-learns
// @Accept  json
// @Produce  json
// @Success 200 {string} string  "ok"
// @Router /lunch_n_learn [get]
func Route(rg *gin.RouterGroup) {
	rg.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"lunch_n_learn": generateMockLunchNLearns(),
			"count":         2,
		})
	})
}

func generateMockLunchNLearns() []lunchnlearn.LunchNLearn {
	mockLunchNLearns := make([]lunchnlearn.LunchNLearn, 2)

	for i := range mockLunchNLearns {
		mockLunchNLearns[i] = lunchnlearn.New("Mocked lunch n learn "+strconv.Itoa(i), nil, []uuid.UUID{uuid.New()}, "2021-01-01 15:04:05")
	}

	return mockLunchNLearns
}
