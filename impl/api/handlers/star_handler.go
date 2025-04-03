package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lrbell17/astroapi/impl/api/dto/request"
	"github.com/lrbell17/astroapi/impl/api/services"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": services.ErrInvalidId})
		return
	}

	star, err := h.service.GetById(uint(id))
	if err == gorm.ErrRecordNotFound {
		log.Warnf("Star with id %d not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		log.Errorf("Unable to get star: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": services.ErrInternal})
		return
	}

	c.JSON(http.StatusOK, star)
}

// Handler function to add star
func (h *StarHandler) Post(c *gin.Context) {

	var req request.StarRequestDTO
	if err := request.ApplyJsonValues(&req, c.Request.Body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	starResp, err := h.service.AddStar(&req)
	if err != nil {

		if services.IsDuplicateKeyErr(err) {
			c.JSON(http.StatusConflict, gin.H{"error": services.ErrStarExists})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": services.ErrInternal})
		return
	}
	c.JSON(http.StatusCreated, starResp)
}

// Handler function to search for stars by name
func (h *StarHandler) SearchByName(c *gin.Context) {
	search, limit := c.Query("search"), c.Query("limit")
	if search == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": services.ErrNoSearchQuery})
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": services.ErrInvalidLimit})
		return
	}

	result, err := h.service.SearchByName(search, limitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": services.ErrInternal})
	}

	c.JSON(http.StatusOK, result)
}
