package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Version string `json:"version"`
}

func HealthHandler(c *gin.Context) {
	var response HealthResponse
	response.Status = "healthy"
	response.Service = "cab-share-backend"
	response.Version = "0.1.0"
	c.JSON(http.StatusOK, response)
}
