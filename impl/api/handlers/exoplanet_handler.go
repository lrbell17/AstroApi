package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lrbell17/astroapi/impl/api/services"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Handler for exoplanet API
type ExoplanetHandler struct {
	service *services.ExoplanetService
}

// Constructor for exoplanet handler
func NewExoplanetHandler(service *services.ExoplanetService) *ExoplanetHandler {
	return &ExoplanetHandler{service}
}

// Handler function for endpoint to get planet by ID
func (h *ExoplanetHandler) GetById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	planet, err := h.service.GetById(uint(id))
	if err == gorm.ErrRecordNotFound {
		log.Warnf("Exoplanet with id %d not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		log.Errorf("Unable too get star: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, planet)
}
