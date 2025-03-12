package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lrbell17/astroapi/impl/api/services"
)

type ExoplanetHandler struct {
	service *services.ExoplanetService
}

func NewExoplanetHandler(service *services.ExoplanetService) *ExoplanetHandler {
	return &ExoplanetHandler{service}
}

func (h *ExoplanetHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	planet, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.JSON(http.StatusOK, planet)
}
