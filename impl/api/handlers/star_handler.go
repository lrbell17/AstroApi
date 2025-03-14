package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lrbell17/astroapi/impl/api/repos"
	"github.com/lrbell17/astroapi/impl/api/services"
	"github.com/lrbell17/astroapi/impl/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Handler for star API
type StarHandler struct {
	service services.StarService
}

// Constructor for star handler
func NewStarHandler(service services.StarService) *StarHandler {
	return &StarHandler{service}
}

// Handler function for getting stars by ID
func (h *StarHandler) GetById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	star, err := h.service.GetById(uint(id))
	if err == gorm.ErrRecordNotFound {
		log.Warnf("Star with id %d not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		log.Errorf("Unable too get star: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, star)
}

// Handler function to add star
func (h *StarHandler) PostStar(c *gin.Context) {
	star := new(model.Star)
	if err := c.BindJSON(star); err != nil {
		log.Errorf("Error parsing star from JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing star from JSON"})
		return
	}

	if err := h.service.AddStar(star); err != nil {

		if err.Error() == repos.ErrStarExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusCreated, star)
}
