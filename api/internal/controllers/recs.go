// Package controllers provides transport layer logic for application endpoints.
package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type RecsController struct {
	log         *slog.Logger
	recommender Recommender
}

type Recommender interface {
	Recommend(ctx context.Context) (string, error)
}

// NewRecsController is a constructor for RecsController.
func NewRecsController(
	log *slog.Logger,
	recommender Recommender,
) *RecsController {
	return &RecsController{
		log:         log,
		recommender: recommender,
	}
}

func (con *RecsController) Recommend(c *gin.Context) {
	res, err := con.recommender.Recommend(context.TODO())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("failed to retrieve recommendations: %s. %v", err.Error(), err),
		})
		return
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(res), &jsonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse recommendations"})
		return
	}

	c.JSON(http.StatusOK, jsonData)
}
